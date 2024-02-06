// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"math/big"
	"testing"
)

func TestPrecalculateNextPortionBlue(t *testing.T) {
	// given a specific portion of stake that is actually blue/red, I want to set preferences of the byzantine nodes
	// to reflect as close a s possible to a portion that
	// map current blue -> expected blue
	// given there are b portion byzantine nodes, and we know the preferences of
	res := precalculateExpectedBluePortion(big.NewFloat(0.1), big.NewFloat(0.9), big.NewFloat(0.001), big.NewFloat(0), 20, 11)
	t.Fatal(res)
}
