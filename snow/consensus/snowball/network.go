// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bag"
	"github.com/ava-labs/avalanchego/utils/sampler"
)

type NewConsensusFunc func(cf ConsensusFactory, params Parameters, choice ids.ID) Consensus

type Network struct {
	params    Parameters
	colors    []ids.ID
	rngSource sampler.Source
	nodes     []Consensus
	virtuous  []Consensus
	running   []Consensus
	cf        ConsensusFactory
}

// Create a new network with [numColors] different possible colors to finalize.
func NewNetwork(cf ConsensusFactory, params Parameters, numColors int, rngSource sampler.Source) *Network {
	n := &Network{
		params:    params,
		rngSource: rngSource,
		cf:        cf,
	}
	for i := 0; i < numColors; i++ {
		n.colors = append(n.colors, ids.Empty.Prefix(uint64(i)))
	}
	return n
}

func (n *Network) AddNode(newConsensusFunc NewConsensusFunc) Consensus {
	s := sampler.NewDeterministicUniform(n.rngSource)
	s.Initialize(uint64(len(n.colors)))
	indices, _ := s.Sample(len(n.colors))
	initialColor := n.colors[int(indices[0])]
	options := make([]int, len(indices)-1)
	for i, option := range indices[1:] {
		options[i] = int(option)
	}
	return n.addNodeWithColor(newConsensusFunc, initialColor, options)
}

// AddNodeSpecificColor adds a new consensus instance to the network which will
// initially prefer [initialPreference] and additionally adds each of the
// specified [options] to consensus.
func (n *Network) AddNodeSpecificColor(
	newConsensusFunc NewConsensusFunc,
	initialPreference int,
	options []int,
) Consensus {
	return n.addNodeWithColor(newConsensusFunc, n.colors[initialPreference], options)
}

func (n *Network) GetColor(index int) ids.ID {
	return n.colors[index]
}

func (n *Network) addNodeWithColor(newConsensusFunc NewConsensusFunc, initialColor ids.ID, options []int) Consensus {
	consensus := newConsensusFunc(n.cf, n.params, initialColor)

	for _, i := range options {
		consensus.Add(n.colors[i])
	}

	n.nodes = append(n.nodes, consensus)
	if !consensus.Finalized() {
		n.running = append(n.running, consensus)
	}

	if byz, ok := consensus.(Byzantiner); !ok {
		n.virtuous = append(n.virtuous, consensus)
	} else {
		byz.SetNetwork(n)
	}

	return consensus
}

// Finalized returns true iff every node added to the network has finished
// running.
func (n *Network) Finalized() bool {
	return len(n.running) == 0
}

// Round simulates a round of consensus by randomly selecting a running node and
// performing an unbiased poll of the nodes in the network for that node.
func (n *Network) Round() {
	if len(n.running) > 0 {
		s := sampler.NewDeterministicUniform(n.rngSource)

		s.Initialize(uint64(len(n.running)))
		runningInd, _ := s.Next()
		running := n.running[runningInd]

		s.Initialize(uint64(len(n.nodes)))
		count := min(n.params.K, len(n.nodes))
		indices, _ := s.Sample(count)
		sampledColors := bag.Bag[ids.ID]{}
		for _, index := range indices {
			peer := n.nodes[int(index)]
			sampledColors.Add(peer.Preference())
		}

		running.RecordPoll(sampledColors)

		// If this node has been finalized, remove it from the poller
		if running.Finalized() {
			newSize := len(n.running) - 1
			n.running[runningInd] = n.running[newSize]
			n.running = n.running[:newSize]
		}
	}
}

// SyncRound simulates a round of consensus where every node sends an unbiased poll
// of the nodes in the network in the current state.
// Nodes only execute their polls AFTER all results have been delivered.
func (n *Network) SyncRound() {
	s := sampler.NewDeterministicUniform(n.rngSource)

	s.Initialize(uint64(len(n.nodes)))
	count := min(n.params.K, len(n.nodes))

	pollResults := make([]bag.Bag[ids.ID], len(n.running))

	for i := range n.running {
		indices, _ := s.Sample(count)
		sampledColors := bag.Bag[ids.ID]{}
		for _, index := range indices {
			peer := n.nodes[int(index)]
			sampledColors.Add(peer.Preference())
		}

		pollResults[i] = sampledColors
	}

	removeIndices := make([]int, 0)
	for i, node := range n.running {
		node.RecordPoll(pollResults[i])

		if node.Finalized() {
			removeIndices = append(removeIndices, i)
		}
	}

	for i, index := range removeIndices {
		newSize := len(n.running) - 1
		n.running[index-i] = n.running[newSize]
		n.running = n.running[:newSize]
	}
}

// Preferences returns the total weight of virtuous nodes behind each preference
// in a slice in the same order as n.colors
func (n *Network) Preferences() []int {
	preferences := make(map[ids.ID]int)
	for _, node := range n.virtuous {
		preferences[node.Preference()]++
	}

	weights := make([]int, len(n.colors))
	for i, color := range n.colors {
		weights[i] = preferences[color]
	}
	return weights
}

// Disagreement returns true iff there are any two nodes in the network that
// have finalized two different preferences.
func (n *Network) Disagreement() bool {
	// Iterate [i] to the index of the first node that has finalized.
	i := 0
	for ; i < len(n.virtuous) && !n.virtuous[i].Finalized(); i++ {
	}
	// If none of the nodes have finalized, then there is no disagreement.
	if i >= len(n.virtuous) {
		return false
	}

	// Return true if any other finalized node has finalized a different
	// preference.
	pref := n.virtuous[i].Preference()
	for ; i < len(n.virtuous); i++ {
		if node := n.virtuous[i]; node.Finalized() && pref != node.Preference() {
			return true
		}
	}
	return false
}

// Agreement returns true iff every node in the network prefers the same value.
func (n *Network) Agreement() bool {
	if len(n.virtuous) == 0 {
		return true
	}
	pref := n.virtuous[0].Preference()
	for _, node := range n.virtuous {
		if pref != node.Preference() {
			return false
		}
	}
	return true
}
