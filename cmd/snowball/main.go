// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"gonum.org/v1/gonum/mathext/prng"
)

// TODO
// switch to synchronous rounds
// add flag for number of simulation runs
// add flag/config to run a set of simulations and output a csv/graph
// switch to target expected value byzantine strategy

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

	network := snowball.NewNetwork(cf, params, 2, source)

	n := v.GetInt(NKey)
	byzantineNodes := int(v.GetFloat64(ByzKey) * float64(n))
	virtuousNodes := n - byzantineNodes
	blueNodes := int(v.GetFloat64(BlueKey) * float64(virtuousNodes))
	redNodes := virtuousNodes - blueNodes

	for i := 0; i < blueNodes; i++ {
		_ = network.AddNodeSpecificColor(snowball.NewFlat, 0, []int{1})
	}
	for i := 0; i < redNodes; i++ {
		_ = network.AddNodeSpecificColor(snowball.NewFlat, 1, []int{0})
	}

	for i := 0; i < byzantineNodes; i++ {
		_ = network.AddNode(snowball.NewByzantine)
	}

	maxRounds := 100_000
	round := 0
	for ; round < maxRounds && !network.Finalized(); round++ {
		network.SyncRound()
	}

	if !network.Finalized() {
		return fmt.Errorf("failed to finalize afer %d rounds", maxRounds)
	}
	if network.Disagreement() {
		return fmt.Errorf("encountered disagreement after %d rounds", round)
	}
	log.Info("Consensus terminated.", zap.Int("rounds", round))
	return nil
}
