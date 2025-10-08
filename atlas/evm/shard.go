// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var (
	_               shard.Shard        = (*vmShard)(nil)
	EVMShardFactory shard.ShardFactory = &vmShardFactory{}
)

type vmShardFactory struct{}

func (f *vmShardFactory) New(ctx context.Context, log logging.Logger, stateDir string) (shard.Shard, error) {
	return New(ctx, log, stateDir)
}

type vmShard struct {
	vm      block.ChainVM
	closeDB func() error
}

func newAdapter(vm block.ChainVM) *vmShard {
	shard := &vmShard{vm: vm}
	return shard
}

func New(
	ctx context.Context,
	log logging.Logger,
	currentStateDir string,
) (*vmShard, error) {
	params, close, err := newVMParams(log, currentStateDir, configBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM params: %w", err)
	}
	vm, err := newFromParams(ctx, params)
	if err != nil {
		close()
		return nil, fmt.Errorf("failed to create VM: %w", err)
	}
	return &vmShard{vm: vm, closeDB: close}, nil
}

func (v *vmShard) CreateHandlers(ctx context.Context) (map[string]http.Handler, error) {
	return v.vm.CreateHandlers(ctx)
}

func (v *vmShard) ExecuteBlock(ctx context.Context, blockBytes []byte) error {
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
	return nil
}

func (v *vmShard) Shutdown(ctx context.Context) error {
	return errors.Join(
		v.vm.Shutdown(ctx),
		v.closeDB(),
	)
}
