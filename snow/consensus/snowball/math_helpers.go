// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"math"
	"math/big"

	"golang.org/x/exp/slices"
)

const precision = 1024

// P(Sum(X_{1-k}) >= alpha | P(X_i = 1) = p ^ P(X_i = 0) = 1-p)
func binomial(p *big.Float, k, alpha int) *big.Float {
	float1 := big.NewFloat(1)
	bigNotP := big.NewFloat(1).SetPrec(precision)
	bigNotP.Sub(bigNotP, p)

	// Exit early if p is 100% or 0
	switch {
	case p.Cmp(float1) == 0:
		return float1
	case bigNotP.Cmp(float1) == 0:
		return big.NewFloat(0)
	default:
	}

	var (
		probability              = big.NewFloat(0).SetPrec(precision)
		pToAlpha                 = big.NewFloat(1).SetPrec(precision)
		notPToNotAlpha           = big.NewFloat(1).SetPrec(precision)
		binomialCoefficientInt   = new(big.Int)
		binomialCoefficientFloat = new(big.Float).SetPrec(precision)
	)
	for i := 0; i < alpha; i++ {
		pToAlpha.Mul(pToAlpha, p)
	}
	for i := alpha; i < k; i++ {
		notPToNotAlpha.Mul(notPToNotAlpha, bigNotP)
	}
	for ; alpha <= k; alpha++ {
		binomialCoefficientFloat.SetInt(binomialCoefficientInt.Binomial(int64(k), int64(alpha))) // k choose alpha
		binomialCoefficientFloat.Mul(binomialCoefficientFloat, pToAlpha).Mul(binomialCoefficientFloat, notPToNotAlpha)
		probability.Add(probability, binomialCoefficientFloat)

		pToAlpha.Mul(pToAlpha, p)
		notPToNotAlpha.Quo(notPToNotAlpha, bigNotP)
	}
	return probability
}

// portionVirtuousBlue is the portion of nodes in the network that are both virtuous and currently blue
// portionByzantine is the portion of nodes that are byzantine
// portionByzantineBlue is the portion of byzantine nodes that are showing blue
func calculateExpectedBluePortionWithByzantine(portionVirtuousBlue, portionByzantine, portionByzantineBlue *big.Float, k, alpha int) *big.Float {
	// Calculate the portion of nodes that are virtuous and will follow the protocol.
	virtuous := big.NewFloat(1)
	virtuous = virtuous.Sub(virtuous, portionByzantine)

	// Calculate the portion of byzantine nodes that are showing blue
	portionByzantineShowingBlue := new(big.Float).Mul(portionByzantine, portionByzantineBlue)

	// Calculate the portion of nodes that will be showing red/blue to calculate flip probabilities
	portionShowingBlue := new(big.Float).Add(portionVirtuousBlue, portionByzantineShowingBlue)
	portionShowingRed := big.NewFloat(1)
	portionShowingRed.Sub(portionShowingRed, portionShowingBlue)

	// Calculate flip probabilities (uniform for the whole network)
	prFlipBlue := binomial(portionShowingBlue, k, alpha)
	prFlipRed := binomial(portionShowingRed, k, alpha)
	prNotFlipRed := big.NewFloat(1)
	prNotFlipRed.Sub(prNotFlipRed, prFlipRed)

	expectedVirtuousStaysBlue := new(big.Float).Mul(portionVirtuousBlue, prNotFlipRed)

	portionVirtuousRed := new(big.Float).Sub(virtuous, portionVirtuousBlue)
	expectedVirtuousFlipBlue := new(big.Float).Mul(portionVirtuousRed, prFlipBlue)

	return new(big.Float).Add(expectedVirtuousStaysBlue, expectedVirtuousFlipBlue)
}

type byzantineVoter struct {
	start, end, interval, portionByzantine float64

	precalculatedResults [][]float64

	blueTarget float64
	k, alpha   int
}

func newByzantineVoter(interval, portionByzantine float64, k, alpha int) *byzantineVoter {
	precalculatedResults := precalculateExpectedBluePortionWithByzantine(0, 1, interval, portionByzantine, k, alpha)

	return &byzantineVoter{
		start:                0, // TODO: remove these as unnecessary if they're not going to be parameterizable
		end:                  1,
		interval:             interval,
		portionByzantine:     portionByzantine,
		precalculatedResults: precalculatedResults,
		blueTarget:           (float64(1) - portionByzantine) / 2,
		k:                    k,
		alpha:                alpha,
	}
}

func (b byzantineVoter) indexAndExpectedWeight(virtuousBlue float64) (int, float64) {
	index := int(math.Round((virtuousBlue - b.start) / b.interval))

	precalculatedResults := b.precalculatedResults[index]
	allRedRes := precalculatedResults[0]
	if allRedRes > b.blueTarget {
		return 0, allRedRes
	}

	allBlueRes := precalculatedResults[len(precalculatedResults)-1]
	if allBlueRes < b.blueTarget {
		return len(precalculatedResults) - 1, allBlueRes
	}

	byzBlueIndex, _ := slices.BinarySearch(precalculatedResults, b.blueTarget)
	return byzBlueIndex, precalculatedResults[byzBlueIndex]
}

func (b byzantineVoter) getByzantinePercentageBlue(virtuousBlue float64) float64 {
	index, _ := b.indexAndExpectedWeight(virtuousBlue)
	return float64(index) * b.interval
}

func precalculateExpectedBluePortionWithByzantine(start, end, interval, portionByzantine float64, k, alpha int) [][]float64 {
	results := make([][]float64, 0)

	// Range from specified start to end
	for portionVirtuousBlue := start; portionVirtuousBlue <= end; portionVirtuousBlue += interval {
		// Range from placing 0 to 100% of byzantine voting power towards blue.
		byzantineBlueToExpectedBlue := make([]float64, 0)
		// This is my bug, this should result in (1 / interval) and it should be (1 / interval) + 1 sine it's inclusive...
		for i := 0; i <= int((end-start)/interval); i += 1 {
			portionByzantineBlue := start + float64(i)*interval
			expectedBlue := calculateExpectedBluePortionWithByzantine(big.NewFloat(portionVirtuousBlue), big.NewFloat(portionByzantine), big.NewFloat(portionByzantineBlue), k, alpha)
			expectedBlueFloat64, _ := expectedBlue.Float64()
			byzantineBlueToExpectedBlue = append(byzantineBlueToExpectedBlue, expectedBlueFloat64)
		}
		results = append(results, byzantineBlueToExpectedBlue)
	}

	return results
}
