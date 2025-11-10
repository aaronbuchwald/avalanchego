// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package shard

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/libevm/log"
	"go.uber.org/zap"
)

type BlockResultHandler interface {
	HandleBlockResult(ctx context.Context, blockResult BlockResult) error
}

func IngestBlockStream(ctx context.Context, blockResultHandler BlockResultHandler, blockResults <-chan BlockResult) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case blockRes, ok := <-blockResults:
			if !ok {
				return errors.New("block results channel closed unexpectedly")
			}
			if err := blockResultHandler.HandleBlockResult(ctx, blockRes); err != nil {
				return fmt.Errorf("failed to execute block %d: %w", blockRes.Height, err)
			}
		}
	}
}

// CreateBlockStreamFromClient creates a channel of block results from the client covering the range (lastAcceptedBlock, tip].
func CreateBlockStreamFromClient(ctx context.Context, client BlockClient, lastAcceptedBlock uint64) (<-chan BlockResult, error) {
	currentBlock := lastAcceptedBlock

	blockResults := make(chan BlockResult, 100)

	go func() {
		defer close(blockResults)

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			latestBlock, err := client.GetMaxHeight(ctx)
			if err != nil {
				log.Error("failed to get max block height",
					zap.Error(err),
					zap.Uint64("currentBlock", currentBlock),
					zap.Uint64("latestBlock", latestBlock),
				)
				return
			}
			// If we're caught up with the tip (current >= latest), wait for new blocks
			if currentBlock >= latestBlock {
				time.Sleep(1 * time.Second)
				continue
			}

			// We're behind the tip, fetch blocks in range (currentBlock, latestBlock]
			err = addBlockRangeToChan(ctx, client, blockResults, currentBlock, latestBlock)
			if err != nil {
				log.Error("failed to add block range to channel",
					zap.Error(err),
					zap.Uint64("currentBlock", currentBlock),
					zap.Uint64("latestBlock", latestBlock),
				)
				return
			}
			currentBlock = latestBlock
		}
	}()

	return blockResults, nil
}

// addBlockRangeToChan fetches blocks in the range (lastAcceptedBlock, latestBlock] and sends them to the channel.
func addBlockRangeToChan(
	ctx context.Context,
	client BlockClient,
	blockResults chan<- BlockResult,
	lastAcceptedBlock uint64,
	latestBlock uint64,
) error {
	for currentBlock := lastAcceptedBlock + 1; currentBlock <= latestBlock; currentBlock++ {
		block, err := client.GetBlockByHeight(ctx, currentBlock)
		if err != nil {
			return fmt.Errorf("failed to get block at height %d: %w", currentBlock, err)
		}
		blockResults <- BlockResult{Height: currentBlock, Block: block}
	}

	return nil
}

type ShardBlockResultHandler struct {
	shard Shard
}

func NewShardBlockResultHandler(shard Shard) *ShardBlockResultHandler {
	return &ShardBlockResultHandler{
		shard: shard,
	}
}

func (h *ShardBlockResultHandler) HandleBlockResult(ctx context.Context, blockResult BlockResult) error {
	return h.shard.ExecuteBlock(ctx, blockResult.Block)
}
