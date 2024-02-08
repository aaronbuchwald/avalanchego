// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/spf13/pflag"
	"gonum.org/v1/gonum/mathext/prng"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("failed due to %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("terminated successfully\n")
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

	switch v.GetString(SnowTypeKey) {
	case "snowball":
		cf = snowball.SnowballFactory{}
	default:
		cf = snowball.SnowflakeFactory{}
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
	i := 0
	for ; i < maxRounds && !network.Finalized(); i++ {
		network.Round()
	}

	if !network.Finalized() {
		return fmt.Errorf("failed to finalize afer %d rounds", maxRounds)
	}
	return nil
}
