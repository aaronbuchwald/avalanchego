// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/stretchr/testify/require"
)

func TestReadShard(t *testing.T) {
	require := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := tests.NewDefaultLogger("test-evm-shard")
	evmShard, err := New(ctx, log, t.TempDir())
	require.NoError(err)
	defer func() {
		require.NoError(evmShard.Shutdown(ctx))
	}()

	shardServer, err := shard.NewServer(ctx, evmShard)
	require.NoError(err)

	testServer := httptest.NewServer(shardServer)
	defer testServer.Close()

	ethClient, err := ethclient.Dial(testServer.URL + "/rpc")
	require.NoError(err)
	defer ethClient.Close()

	blockNumber, err := ethClient.BlockNumber(ctx)
	require.NoError(err)
	require.Equal(blockNumber, uint64(0))
}
