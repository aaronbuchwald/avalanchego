// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockdb

import (
	"context"
	"testing"

	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/stretchr/testify/require"
)

func TestIngestBlockStream(t *testing.T) {
	require := require.New(t)

	blockDB := newTestBlockDB()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a channel and send some blocks
	blockResults := make(chan shard.BlockResult)

	// Create handler using Strategy pattern
	handler := NewBlockDBResultHandler(blockDB)

	// Run IngestBlockStream in a goroutine
	errChan := make(chan error)
	go func() {
		errChan <- shard.IngestBlockStream(ctx, handler, blockResults)
	}()

	// Send blocks
	testBlocks := []shard.BlockResult{
		{Height: 0, Block: []byte("genesis block")},
		{Height: 1, Block: []byte("block 1")},
		{Height: 2, Block: []byte("block 2")},
	}

	for _, block := range testBlocks {
		blockResults <- block
	}

	// Ingest one more block to block until the original testBlocks are ingested
	blockResults <- shard.BlockResult{Height: 3, Block: []byte("block 3")}

	// Verify blocks were written
	for _, expected := range testBlocks {
		actual, err := blockDB.GetBlockByHeight(expected.Height)
		require.NoError(err)
		require.Equal(expected.Block, actual)
	}
	cancel()
	require.NoError(<-errChan)
}
