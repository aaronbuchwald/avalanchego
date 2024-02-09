// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"math"
	"math/big"

	"golang.org/x/exp/slices"
)

const precision = 1024

func binomial(p *big.Float, k, min, max int) *big.Float {
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
	alpha := min
	for i := 0; i < alpha; i++ {
		pToAlpha.Mul(pToAlpha, p)
	}
	for i := alpha; i < k; i++ {
		notPToNotAlpha.Mul(notPToNotAlpha, bigNotP)
	}
	for ; alpha <= max; alpha++ {
		binomialCoefficientFloat.SetInt(binomialCoefficientInt.Binomial(int64(k), int64(alpha))) // k choose alpha
		binomialCoefficientFloat.Mul(binomialCoefficientFloat, pToAlpha).Mul(binomialCoefficientFloat, notPToNotAlpha)
		probability.Add(probability, binomialCoefficientFloat)

		pToAlpha.Mul(pToAlpha, p)
		notPToNotAlpha.Quo(notPToNotAlpha, bigNotP)
	}
	return probability
}

// P(Sum(X_{1-k}) >= alpha | P(X_i = 1) = p ^ P(X_i = 0) = 1-p)
func binomialGE(p *big.Float, k, alpha int) *big.Float {
	return binomial(p, k, alpha, k)
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
	prFlipBlue := binomialGE(portionShowingBlue, k, alpha)
	prFlipRed := binomialGE(portionShowingRed, k, alpha)
	prNotFlipRed := big.NewFloat(1)
	prNotFlipRed.Sub(prNotFlipRed, prFlipRed)

	expectedVirtuousStaysBlue := new(big.Float).Mul(portionVirtuousBlue, prNotFlipRed)

	portionVirtuousRed := new(big.Float).Sub(virtuous, portionVirtuousBlue)
	expectedVirtuousFlipBlue := new(big.Float).Mul(portionVirtuousRed, prFlipBlue)

	return new(big.Float).Add(expectedVirtuousStaysBlue, expectedVirtuousFlipBlue)
}

type ByzantineVoter struct {
	start, end, interval, portionByzantine float64

	precalculatedResults [][]float64

	blueTarget float64
	k, alpha   int
}

func NewByzantineVoter(interval, portionByzantine float64, k, alpha int) *ByzantineVoter {
	precalculatedResults := precalculateExpectedBluePortionWithByzantine(0, 1, interval, portionByzantine, k, alpha)
	return &ByzantineVoter{
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

func (b ByzantineVoter) virtuousBlueResultsIndex(virtuousBlue float64) int {
	return int(math.Round((virtuousBlue - b.start) / b.interval))
}

func (b ByzantineVoter) indexAndExpectedWeight(virtuousBlue float64) (int, float64) {
	index := b.virtuousBlueResultsIndex(virtuousBlue)

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

func (b ByzantineVoter) GetByzantinePercentageBlue(virtuousBlue float64) float64 {
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

func prDecideReachAlpha(p, byzantine float64, k, alpha int) []*big.Float {
	prDecidingVotes := make([]*big.Float, 0)
	// Iterate the range [1, alpha]
	// If I control 0 votes, I have no influence.
	// If I control alpha votes, I decide deterministically.
	for i := 0; i <= alpha; i++ {
		// Calculate the probability the byzantine party controls i votes in a poll of size k
		prIVotes := binomial(big.NewFloat(byzantine), k, i, i)

		// Given p, what is the probability that my i votes cast one way or the other
		// dictates whether an alpha threshold is reached.
		// Assuming that all of the votes that did not go to the byzantine party, went to
		// the virtuous party, there are k - i remaining trials.
		// Assume p is the probability of success (out of the virtuous population)
		// If the number of successes is in the range [alpha-i, alpha-1] then our allotment
		// of votes decides the outcome.
		// Note: if we want to decide the outcome one way or the other, the dominant strategy
		// is always to vote 100% in that direction.
		// However, if the rest of the poll results are unknown, BUT have a known probability
		// distribution, then we can calculate a more accurate expected value by accounting
		// for the number of votes within a single poll.
		prDecidingVote := binomial(big.NewFloat(p), k-i, alpha-i, alpha-1)
		prDecidingVote.Mul(prIVotes, prDecidingVote)

		prDecidingVotes = append(prDecidingVotes, prDecidingVote)
	}

	// Note: if a byzantine adversary wants to minimize the variance of possible outcomes, a simple strategy
	// is to note that it has the tightest control of the outcome the more votes it has.
	// Therefore, if it knows that it should wants to flip a given number of nodes, it should prioritize
	// flipping the nodes where its decision reduces the variance the most.
	// Assuming that the network is split more evenly than the byzantine adversary (safe assumption assuming
	// the network hasn't reached the tipping point), then we reduce the variance the most by either
	// 1. Voting entirely blue in the polls where we are included most
	// 2. Voting entirely blue in the polls where we are included the least
	// Intuitively, it seems like we should vote prioritize voting entirely blue when we have the best chance
	// of influencing the outcome. (TODO: check the math)
	// UPDATE: building on the intuition, the variance is lower if we vote uniformly in polls where we are included
	// the most since the variance is the sum of variances coming from the virtuous portion of the network
	// and the byzantine portion (us) for whom it is uniformly blue ie. Var(byz) = 0. If we prioritized
	// the polls where we had fewer, then we'd potentially be left splitting the polls where we have the
	// most control and the ability to reduce the variance the most.
	// Q: does it matter that this results in an uneven distribution ie. not a perfect bell curve centered
	// around our target expected value. It may reduce the variance, but increase the frequency of tail events.
	// As a result, given that this is the only information I have access to, I argue this is in fact the
	// optimal strategy.
	// Further, for a network of size n >= 500, the variance should be extremely low, so the choice
	// of whether prioritizing larger or smaller polls to reduce the distribution of outcomes should
	// have only a small impact.
	// TODO: can I make a simple convincing argument/proof that minimizing the variance in this way is
	// optimal?
	return prDecidingVotes
}

// In addition, we probably also care about how many nodes that we have no influence over are expected
// to end up being blue/red.
// Given that, we can calculate the exact range where in expectation we can push the node back to 50-50.
func prBlueWithNoByzantineInfluence(p, byzantine float64, k, alpha int) float64 {
	virtuous := big.NewFloat(1)
	virtuous.Sub(virtuous, big.NewFloat(byzantine))
	prNoByzantineVotes := binomialGE(virtuous, k, k)

	prNoByzantineBlue := binomialGE(big.NewFloat(p), k, alpha)
	prNoByzantineBlue.Mul(prNoByzantineBlue, prNoByzantineVotes)
	fl, _ := prNoByzantineBlue.Float64()
	return fl
}
