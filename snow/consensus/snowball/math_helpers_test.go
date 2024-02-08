// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"fmt"
	"strings"
	"testing"
)

func TestByzantineVoter(t *testing.T) {
	var (
		interval = float64(0.01)
		byz      = float64(0.02)
	)
	bv := newByzantineVoter(interval, byz, 20, 11)

	sb := strings.Builder{}
	for virtuousBlue, byzBlueToExpectedBlue := range bv.precalculatedResults {
		var (
			start = 0
			end   = len(byzBlueToExpectedBlue) - 1
		)
		allRedRes := byzBlueToExpectedBlue[start]
		// if answering all red, still puts us above the blue target range, then vote minority strategy dominates
		if allRedRes > bv.blueTarget {
			continue
		}

		// if answering all blue, still puts us below the blue target range, then vote minority strategy dominates
		allBlueRes := byzBlueToExpectedBlue[end]
		if allBlueRes < bv.blueTarget {
			continue
		}

		// if the minority strategy does not dominate, then output the interesting section of the results
		sb.WriteString(fmt.Sprintf("\n%.2f", float64(virtuousBlue)*interval))

		for byzBlueIndex, expectedBlue := range byzBlueToExpectedBlue {
			byzBlue := float64(byzBlueIndex) * interval
			sb.WriteString(fmt.Sprintf("\n(ByzBlue = %.2f, ExpectedBlue = %.2f)", byzBlue, expectedBlue))
		}
	}
	t.Fatal(sb.String())
}
