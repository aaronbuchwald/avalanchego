// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import "math/big"

const precision = 1024

// P(Sum(X_{1-k}) >= alpha | P(X_i = 1) = p ^ P(X_i = 0) = 1-p)
func binomial(p *big.Float, k, alpha int) *big.Float {
	bigNotP := big.NewFloat(1).SetPrec(precision)
	bigNotP.Sub(bigNotP, p)

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
	virtuous := big.NewFloat(1)
	virtuous = virtuous.Sub(virtuous, portionByzantine)

	portionByzantineShowingBlue := new(big.Float).Mul(portionByzantine, portionByzantineBlue)

	portionShowingBlue := new(big.Float).Add(portionVirtuousBlue, portionByzantineShowingBlue)
	portionShowingRed := big.NewFloat(1)
	portionShowingRed.Sub(portionShowingRed, portionShowingBlue)

	prFlipBlue := binomial(portionShowingBlue, k, alpha)
	prFlipRed := binomial(portionShowingRed, k, alpha)
	prNotFlipRed := big.NewFloat(1)
	prNotFlipRed.Sub(prNotFlipRed, prFlipRed)

	expectedVirtuousStaysBlue := new(big.Float).Mul(portionVirtuousBlue, prNotFlipRed)

	portionVirtuousRed := new(big.Float).Sub(virtuous, portionVirtuousBlue)
	expectedVirtuousFlipBlue := new(big.Float).Mul(portionVirtuousRed, prFlipBlue)

	return new(big.Float).Add(expectedVirtuousStaysBlue, expectedVirtuousFlipBlue)

}

// calculateExpectedBluePortion calculates the expected portion of nodes that will be blue in the next round
// given all nodes follow the protocol with the given parameters.
func calculateExpectedBluePortion(portionBlue, portionByzantine *big.Float, k, alpha int) *big.Float {
	portionRed := big.NewFloat(1)
	portionRed.Sub(portionRed, portionBlue)

	prFlipRed := binomial(portionRed, k, alpha)
	expectedPortionFlipRed := new(big.Float).Mul(portionBlue, prFlipRed)
	expectedPortionStayBlue := new(big.Float).Sub(portionBlue, expectedPortionFlipRed)

	prFlipBlue := binomial(portionBlue, k, alpha)
	expectedPortionFlipBlue := new(big.Float).Mul(portionRed, prFlipBlue)

	return new(big.Float).Add(expectedPortionStayBlue, expectedPortionFlipBlue)
}

func precalculateExpectedBluePortion(start, end, interval, portionByzantine *big.Float, k, alpha int) map[float64]float64 {
	res := make(map[float64]float64)
	for portionBlue := start; portionBlue.Cmp(end) <= 0; portionBlue.Add(portionBlue, interval) {
		portionBlueFloat, _ := portionBlue.Float64()
		expectedPortionBlue, _ := calculateExpectedBluePortion(portionBlue, portionByzantine, k, alpha).Float64()
		res[portionBlueFloat] = expectedPortionBlue
	}

	return res
}
