// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"fmt"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/sampler"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Simulation struct {
	log            logging.Logger
	network        *Network
	byzantine      []*Byzantine
	byzantineVoter *ByzantineVoter
	stepThrough    bool
}

func NewSimulation(
	log logging.Logger,
	cf Factory,
	params Parameters,
	numNodes int,
	initialVirtuousBlue float64,
	byz float64,
	rngSource sampler.Source,
	stepThrough bool,
) *Simulation {
	network := NewNetwork(cf, params, 2, rngSource)
	byzantineNodes := int(float64(numNodes) * byz)
	virtuousNodes := numNodes - byzantineNodes
	blueNodes := int(float64(virtuousNodes) * initialVirtuousBlue)
	redNodes := virtuousNodes - blueNodes

	byzVoter := NewByzantineVoter(0.01, byz, params.K, params.AlphaPreference)
	s := &Simulation{
		network:        network,
		byzantineVoter: byzVoter,
		log:            log,
		stepThrough:    stepThrough,
	}

	for i := 0; i < blueNodes; i++ {
		_ = network.AddNodeSpecificColor(NewFlat, 0, []int{1})
	}
	for i := 0; i < redNodes; i++ {
		_ = network.AddNodeSpecificColor(NewFlat, 1, []int{0})
	}
	for i := 0; i < byzantineNodes; i++ {
		s.byzantine = append(s.byzantine, network.AddNode(NewByzantine).(*Byzantine))
	}
	return s
}

func (s *Simulation) Execute(
	maxRounds int,
) (int, error) {
	round := 0
	blue := s.network.colors[0]
	red := s.network.colors[1]
	for ; round < maxRounds && !s.network.Finalized(); round++ {
		preferences := s.network.Preferences()

		bluePreference := preferences[0]
		virtuousBlue := float64(bluePreference) / float64(len(s.network.nodes)-len(s.byzantine))
		blueOverN := float64(bluePreference) / float64(len(s.network.nodes))
		byzShowBluePortion, expectedNextBlue := s.byzantineVoter.GetByzantinePercentageBlueAndExpectedBlue(float64(bluePreference) / float64(len(s.network.nodes)))
		blueCutoffPoint := int(byzShowBluePortion * float64(len(s.byzantine)))
		for i, byz := range s.byzantine {
			if i < blueCutoffPoint {
				byz.SetPreference(blue)
			} else {
				byz.SetPreference(red)
			}
		}
		portionShowingBlue := float64(blueCutoffPoint+bluePreference) / float64(len(s.network.nodes))
		s.log.Info("Sim step",
			zap.Int("round", round),
			zap.Float64("blue/virtuous", virtuousBlue),
			zap.Float64("blue/n", blueOverN),
			zap.Float64("byzShowingBlue", byzShowBluePortion),
			zap.Float64("portionShowingBlue", portionShowingBlue),
			zap.Float64("expectedNextBlue", expectedNextBlue),
		)

		if s.stepThrough {
			fmt.Printf("Hit enter to continue:\n")
			fmt.Scanln()
		}

		s.network.SyncRound()
	}

	if !s.network.Finalized() {
		return 0, fmt.Errorf("failed to finalize after %d rounds", maxRounds)
	}
	if s.network.Disagreement() {
		return 0, fmt.Errorf("found disagreement after %d rounds", round)
	}

	return round, nil
}

func ExecuteSimulations(
	log logging.Logger,
	cf Factory,
	params Parameters,
	numNodes int,
	initialVirtuousBlue float64,
	byz float64,
	rngSource sampler.Source,
	numSimulations int,
	maxRounds int,
	stepThrough bool,
) []int {
	eg := errgroup.Group{}
	eg.SetLimit(10)
	simulationResults := make([]int, numSimulations)
	for i := 0; i < numSimulations; i++ {
		i := i
		eg.Go(func() error {

			sim := NewSimulation(log, cf, params, numNodes, initialVirtuousBlue, byz, rngSource, stepThrough)
			roundsToTermination, err := sim.Execute(maxRounds)
			if err != nil {
				log.Warn("simulation failed", zap.Error(err))
				simulationResults[i] = maxRounds // Record maxRounds in place of failure
				return nil
			}

			simulationResults[i] = roundsToTermination
			return nil
		})
	}
	eg.Wait()

	return simulationResults
}

func ExecuteSimulationsWithDifferentSizeByzantineAdversaries(
	log logging.Logger,
	cf Factory,
	params Parameters,
	numNodes int,
	initialVirtuousBlue float64,
	byzAdversaries []float64,
	rngSource sampler.Source,
	numSimulations int,
	maxRounds int,
	stepThrough bool,
) map[float64][]int {
	byzResults := make(map[float64][]int)
	for _, byz := range byzAdversaries {
		res := ExecuteSimulations(
			log,
			cf,
			params,
			numNodes,
			initialVirtuousBlue,
			byz,
			rngSource,
			numSimulations,
			maxRounds,
			stepThrough,
		)

		byzResults[byz] = res
	}

	return byzResults
}
