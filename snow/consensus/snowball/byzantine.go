// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bag"
)

var (
	Red   = ids.Empty.Prefix(0)
	Blue  = ids.Empty.Prefix(1)
	Green = ids.Empty.Prefix(2)

	_ Byzantiner = (*Byzantine)(nil)
)

type Byzantiner interface {
	Consensus
	SetNetwork(*Network)
}

func NewByzantine(_ Factory, _ Parameters, choice ids.ID) Consensus {
	return &Byzantine{
		preference: choice,
	}
}

// Byzantine is a naive implementation of a multi-choice snowball instance
type Byzantine struct {
	// Hardcode the preference
	preference ids.ID
}

func (*Byzantine) Add(ids.ID) {}

func (b *Byzantine) Preference() ids.ID {
	return b.preference
}

func (b *Byzantine) SetPreference(pref ids.ID) {
	b.preference = pref
}

func (*Byzantine) RecordPoll(bag.Bag[ids.ID]) bool {
	return false
}

func (*Byzantine) RecordUnsuccessfulPoll() {}

func (*Byzantine) Finalized() bool {
	return true
}

func (b *Byzantine) String() string {
	return b.preference.String()
}

func (b *Byzantine) SetNetwork(*Network) {}

func NewByzantineMinorityVote(_ Factory, _ Parameters, _ ids.ID) Consensus {
	return &ByzantineMinorityVote{}
}

// ByzantineMinorityVote is a byzantine implementation of consensus that always
// votes for the current minority color of the network.
type ByzantineMinorityVote struct {
	network *Network
}

func (*ByzantineMinorityVote) Add(ids.ID) {}

func (b *ByzantineMinorityVote) Preference() ids.ID {
	// TODO: switch to maintaining preference count inside the network to avoid reapting
	// calculation for each byzantine node.
	preferences := make(map[ids.ID]int)
	for _, node := range b.network.virtuous {
		preferences[node.Preference()]++
	}

	minColor := b.network.colors[0]
	minColorPref := preferences[minColor]

	for color, pref := range preferences {
		if pref < minColorPref {
			minColor = color
			minColorPref = pref
		}
	}
	return minColor
}

func (*ByzantineMinorityVote) RecordPoll(bag.Bag[ids.ID]) bool {
	return false
}

func (*ByzantineMinorityVote) RecordUnsuccessfulPoll() {}

func (*ByzantineMinorityVote) Finalized() bool {
	return true
}

func (b *ByzantineMinorityVote) String() string {
	return "ByzantineMinorityVoter"
}

func (b *ByzantineMinorityVote) SetNetwork(n *Network) {
	b.network = n
}
