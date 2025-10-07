// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package shard

import (
	"context"
	"net/http"

	"github.com/ava-labs/avalanchego/api/health"
)

type ReadShard interface {
	health.Checker

	CreateHandlers(ctx context.Context) (map[string]http.Handler, error)
	GetAvailableRange() (uint64, uint64, error)
}

type WriteShard interface {
	ExecuteBlock(ctx context.Context, blockBytes []byte) error
}

type Shard interface {
	ReadShard
	WriteShard
}
