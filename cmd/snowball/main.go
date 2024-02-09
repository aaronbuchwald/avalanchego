// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"gonum.org/v1/gonum/mathext/prng"
)

// TODO
// create a function given a set of parameters to run a single simulation and output rounds to termination / failure
// separate out the functionality to output a CSV file
// create graphs of 1) scatter plot of rounds to termination with a given byz % 2) byz % vs distribution of rounds to termination over n simulations
// debug byz strategy with alpha pref set to k/2 + 1

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("failed due to %s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	v, err := BuildViper(args)
	if errors.Is(err, pflag.ErrHelp) {
		os.Exit(0)
	}
	if err != nil {
		fmt.Printf("failed to build config: %s\n", err)
		os.Exit(1)
	}

	level, err := logging.ToLevel(v.GetString(LogLevelKey))
	if err != nil {
		return err
	}
	log := logging.NewLogger(
		"snowball-simulation",
		logging.NewWrappedCore(
			level,
			os.Stdout,
			logging.Colors.ConsoleEncoder(),
		),
	)

	var (
		cf     snowball.ConsensusFactory
		params = snowball.Parameters{
			K:               v.GetInt(KKey),
			AlphaConfidence: v.GetInt(AlphaConfidenceKey),
			AlphaPreference: v.GetInt(AlphaPreferenceKey),
			BetaVirtuous:    v.GetInt(BetaVirtuousKey),
			BetaRogue:       v.GetInt(BetaRogueKey),
		}
		seed   uint64 = 0
		source        = prng.NewMT19937()
	)
	source.Seed(seed)

	switch snowType := v.GetString(SnowTypeKey); snowType {
	case "snowball":
		cf = snowball.SnowballFactory{}
	case "snowflake":
		cf = snowball.SnowflakeFactory{}
	default:
		return fmt.Errorf("invalid snow type input: %q", snowType)
	}

	writer := io.Writer(log)
	outputPath := v.GetString(OutputFileKey)
	if len(outputPath) != 0 {
		f, err := os.Create(os.ExpandEnv(outputPath))
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()
		writer = io.MultiWriter(writer, f)
	}
	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	csvWriter.Write([]string{"sim", "rounds"})

	stepThrough := v.GetBool(StepThroughKey)
	maxRounds := v.GetInt(MaxRoundsKey)
	numSimulations := v.GetInt(NumSimulationsKey)
	n := v.GetInt(NKey)
	byzantinePortion := v.GetFloat64(ByzKey)
	numByzantineNodes := int(byzantinePortion * float64(n))
	virtuousNodes := n - numByzantineNodes
	blueNodes := int(v.GetFloat64(BlueKey) * float64(virtuousNodes))
	redNodes := virtuousNodes - blueNodes

	byzVoter := snowball.NewByzantineVoter(0.01, byzantinePortion, params.K, params.AlphaPreference)

	for sim := 0; sim < numSimulations; sim++ {
		network := snowball.NewNetwork(cf, params, 2, source)
		blue := network.GetColor(0)
		red := network.GetColor(1)

		for i := 0; i < blueNodes; i++ {
			_ = network.AddNodeSpecificColor(snowball.NewFlat, 0, []int{1})
		}
		for i := 0; i < redNodes; i++ {
			_ = network.AddNodeSpecificColor(snowball.NewFlat, 1, []int{0})
		}

		byzantineNodes := make([]*snowball.Byzantine, numByzantineNodes)
		for i := 0; i < numByzantineNodes; i++ {
			byzantineNodes[i] = network.AddNode(snowball.NewByzantine).(*snowball.Byzantine)
		}

		round := 0
		for ; round < maxRounds && !network.Finalized(); round++ {
			network.SyncRound()

			preferences := network.Preferences()
			blueWeight := preferences[0]
			blueVirtuousPortion := float64(blueWeight) / float64(virtuousNodes)
			blueVirtuousOverN := float64(blueWeight) / float64(n)
			byzBluePortion := byzVoter.GetByzantinePercentageBlue(blueVirtuousOverN)
			blueCutoff := int(byzBluePortion * float64(len(byzantineNodes)))
			totalBlue := blueWeight + blueCutoff

			log.Info("Preference update",
				zap.Int("sim", sim),
				zap.Int("round", round),
				zap.Float64("blueVirtuousPortion", blueVirtuousPortion),
				zap.Float64("blueVirtuousOverN", blueVirtuousOverN),
				zap.Float64("byzBlue", byzBluePortion),
				zap.Int("byzBlueCutoff", blueCutoff),
				zap.Int("byzNodes", numByzantineNodes),
				zap.Int("totalBlue", totalBlue),
			)

			if stepThrough {
				fmt.Printf("press enter to continue\n")
				fmt.Scanln()
			}

			for i, byz := range byzantineNodes {
				if i < blueCutoff {
					byz.SetPreference(blue)
				} else {
					byz.SetPreference(red)
				}
			}
		}

		if !network.Finalized() {
			return fmt.Errorf("failed to finalize afer %d rounds", maxRounds)
		}
		if network.Disagreement() {
			return fmt.Errorf("encountered disagreement after %d rounds", round)
		}

		if err := csvWriter.Write([]string{
			fmt.Sprintf("%d", sim),
			fmt.Sprintf("%d", round),
		}); err != nil {
			return fmt.Errorf("failed to write output on sim %d: %w", sim, err)
		}
	}
	return nil
}
