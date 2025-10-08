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

type ReadShard interface {
	Shutdown(ctx context.Context) error
	CreateHandlers(ctx context.Context) (map[string]http.Handler, error)
	// TODO: add way for shards to advertise available range
}

type WriteShard interface {
	ExecuteBlock(ctx context.Context, blockBytes []byte) error
}

type Shard interface {
	ReadShard
	WriteShard
}
