// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"testing"
)

func TestByzantineVoter(t *testing.T) {
	bv := newByzantineVoter(0.01, 0.1, 20, 11)

	index, expectedNextBlue := bv.indexAndExpectedWeight(0.1)
	byzBlue := bv.getByzantinePercentageBlue(0.1)
	t.Fatal(index, expectedNextBlue, byzBlue, len(bv.precalculatedResults[0]), bv.precalculatedResults)
}
