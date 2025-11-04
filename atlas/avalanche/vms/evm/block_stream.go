// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"math/big"

	"github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/libevm/rlp"
)

type EVMBlockClient struct {
	client *ethclient.Client
}

func NewEVMBlockClient(client *ethclient.Client) *EVMBlockClient {
	return &EVMBlockClient{client}
}

func (c *EVMBlockClient) GetBlockByHeight(ctx context.Context, height uint64) ([]byte, error) {
	block, err := c.client.BlockByNumber(ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	blockBytes, err := rlp.EncodeToBytes(block)
	if err != nil {
		return nil, err
	}
	return blockBytes, nil
}

func (c *EVMBlockClient) GetMaxHeight(ctx context.Context) (uint64, error) {
	return c.client.BlockNumber(ctx)
}
