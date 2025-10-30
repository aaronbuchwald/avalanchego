// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package shard

import (
	"context"
	"net/http"

	"github.com/ava-labs/avalanchego/utils/logging"
)

type ShardFactory interface {
	New(ctx context.Context, log logging.Logger, stateDir string) (Shard, error)
}

type HeightRange struct {
	Start uint64  // Start height included in the range, nil indicates genesis.
	End   *uint64 // End height included in the range, nil indicates tip.
}

type ReadShard interface {
	Shutdown(ctx context.Context) error
	CreateHandlers(ctx context.Context) (map[string]http.Handler, error)
	// TODO: add HeightRange to ReadShard interface
	// HeightRange(ctx context.Context) (HeightRange, error)
}

type BlockResult struct {
	Height uint64
	Block  []byte
}

type WriteShard interface {
	ExecuteBlock(ctx context.Context, blockBytes []byte) error
}

type SplittableShard interface {
	SplitAtHeight(ctx context.Context, targetHeight uint64, targetStateDir string) error
}

type Shard interface {
	ReadShard
	WriteShard
	SplittableShard
}

