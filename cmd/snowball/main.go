// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/spf13/pflag"
	"gonum.org/v1/gonum/mathext/prng"
)

// TODO
// debug byz strategy with alpha pref set to k/2 + 1
// graph the data potentially in python

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("failed due to %s\n", err)
		os.Exit(1)
	}
}

func writeSimulationResultsToCSV(outputPath string, results []int) error {
	f, err := os.Create(os.ExpandEnv(outputPath))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	csvWriter := csv.NewWriter(f)
	defer csvWriter.Flush()

	if err := csvWriter.Write([]string{"sim", "rounds"}); err != nil {
		return err
	}

	for i, result := range results {
		if err := csvWriter.Write([]string{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%d", result),
		}); err != nil {
			return err
		}
	}

	return nil
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

	byzPortionsInt := v.GetIntSlice(ByzKey)
	byzPortions := make([]float64, len(byzPortionsInt))
	for i, byzInt := range byzPortionsInt {
		byzPortions[i] = float64(byzInt) / float64(100)
	}

	byzResults := snowball.ExecuteSimulationsWithDifferentSizeByzantineAdversaries(
		log,
		cf,
		params,
		snowball.NewFlat,
		v.GetInt(NKey),
		v.GetFloat64(BlueKey),
		byzPortions,
		source,
		v.GetInt(NumSimulationsKey),
		v.GetInt(MaxRoundsKey),
	)

	outputFilePath := v.GetString(OutputFileKey)
	for byzPortion, results := range byzResults {
		if len(outputFilePath) > 0 {
			if err := writeSimulationResultsToCSV(fmt.Sprintf("%.2f_%s.csv", byzPortion, outputFilePath), results); err != nil {
				return err
			}
		}
	}
	return nil
}
