/*
Copyright © 2025 Aaron Buchwald <aaron.buchwald56@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/libevm/rlp"
	"go.uber.org/zap"
)

// createBlockResultStreamFromClient creates a channel of block results from the client.
func createBlockResultStreamFromClient(ctx context.Context, client *ethclient.Client, log logging.Logger, lastAcceptedBlock uint64) (<-chan shard.BlockResult, error) {
	currentBlock := lastAcceptedBlock

	blockResults := make(chan shard.BlockResult, 100)

	go func() {
		defer close(blockResults)

		for {
			select {
			case <-ctx.Done():
				log.Error("context done", zap.Error(ctx.Err()))
				return
			default:
			}

			latestBlock, err := client.BlockNumber(ctx)
			if err != nil {
				log.Error("failed to get block number",
					zap.Error(err),
					zap.Uint64("currentBlock", currentBlock),
					zap.Uint64("latestBlock", latestBlock),
				)
				return
			}
			if currentBlock <= latestBlock {
				time.Sleep(1 * time.Second)
				continue
			}

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

// addBlockRangeToChan adds the range of blocks (lastAcceptedBlock, latestBlock] to
// blockResults.
func addBlockRangeToChan(
	ctx context.Context,
	client *ethclient.Client,
	blockResults chan<- shard.BlockResult,
	lastAcceptedBlock uint64,
	latestBlock uint64,
) error {
	for currentBlock := lastAcceptedBlock + 1; currentBlock <= latestBlock; currentBlock++ {
		block, err := client.BlockByNumber(ctx, big.NewInt(int64(currentBlock)))
		if err != nil {
			return fmt.Errorf("failed to get block: %w", err)
		}
		blockBytes, err := rlp.EncodeToBytes(block)
		if err != nil {
			return fmt.Errorf("failed to encode block: %w", err)
		}
		blockResults <- shard.BlockResult{Height: currentBlock, Block: blockBytes}
	}

	return nil
}
