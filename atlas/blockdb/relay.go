// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockdb

import (
	"context"

	"github.com/ava-labs/avalanchego/atlas/shard"
)

type BlockDBResultHandler struct {
	db *BlockDB
}

func NewBlockDBResultHandler(db *BlockDB) *BlockDBResultHandler {
	return &BlockDBResultHandler{
		db: db,
	}
}

func (h *BlockDBResultHandler) HandleBlockResult(ctx context.Context, blockResult shard.BlockResult) error {
	return h.db.WriteBlock(blockResult.Height, blockResult.Block)
}

func IngestBlockStream(ctx context.Context, db *BlockDB, blockResults <-chan shard.BlockResult) error {
	blockResultHandler := NewBlockDBResultHandler(db)
	return shard.IngestBlockStream(ctx, blockResultHandler, blockResults)
}
