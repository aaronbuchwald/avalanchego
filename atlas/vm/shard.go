// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var _ shard.Shard = (*VMShardAdapter)(nil)

type VMShardAdapter struct {
	vm                     block.ChainVM
	startHeight, endHeight atomic.Uint64
}

func NewVMShardAdapter(vm block.ChainVM, startHeight, endHeight uint64) *VMShardAdapter {
	shard := &VMShardAdapter{vm: vm}
	shard.startHeight.Store(startHeight)
	shard.endHeight.Store(endHeight)
	return shard
}

func (v *VMShardAdapter) CreateHandlers(ctx context.Context) (map[string]http.Handler, error) {
	return v.vm.CreateHandlers(ctx)
}

func (v *VMShardAdapter) HealthCheck(ctx context.Context) (interface{}, error) {
	return v.vm.HealthCheck(ctx)
}

func (v *VMShardAdapter) GetAvailableRange() (uint64, uint64, error) {
	return v.startHeight.Load(), v.endHeight.Load(), nil
}

func (v *VMShardAdapter) ExecuteBlock(ctx context.Context, blockBytes []byte) error {
	blk, err := v.vm.ParseBlock(ctx, blockBytes)
	if err != nil {
		return fmt.Errorf("failed to parse block: %w", err)
	}
	if err := blk.Verify(ctx); err != nil {
		return fmt.Errorf("failed to verify block: %w", err)
	}
	if err := blk.Accept(ctx); err != nil {
		return fmt.Errorf("failed to accept block: %w", err)
	}

	v.endHeight.Store(blk.Height())
	return nil
}
