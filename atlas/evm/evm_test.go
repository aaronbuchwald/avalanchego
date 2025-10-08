// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"embed"
	"fmt"
	"io"
	"net"
	"net/http/httptest"
	"testing"

	pb "github.com/ava-labs/avalanchego/atlas/proto/pb/writeshard"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:embed blockdata/*
var blockDataFiles embed.FS

type shardTest struct {
	require *require.Assertions
	ctx     context.Context
	cancel  context.CancelFunc
	shard   *vmShard
}

func setup(tb testing.TB) *shardTest {
	require := require.New(tb)
	ctx, cancel := context.WithCancel(context.Background())

	log := tests.NewDefaultLogger("test-evm-shard")
	var err error
	evmShard, err := New(ctx, log, tb.TempDir())
	require.NoError(err)
	tb.Cleanup(func() {
		cancel()
		require.NoError(evmShard.Shutdown(ctx))
	})

	return &shardTest{
		require: require,
		ctx:     ctx,
		cancel:  cancel,
		shard:   evmShard,
	}
}

func TestReadShard(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shard
	defer cancel()

	shardServer, err := shard.NewServer(shardTest.ctx, vmShard)
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

func TestActiveShard(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shard
	defer cancel()

	shardServer, err := shard.NewServer(shardTest.ctx, vmShard)
	require.NoError(err)

	testServer := httptest.NewServer(shardServer)
	defer testServer.Close()

	listener, err := net.Listen("tcp", ":0")
	require.NoError(err)
	defer listener.Close()

	go shard.ServeGRPCShard(ctx, listener, vmShard)

	ethClient, err := ethclient.Dial(testServer.URL + "/rpc")
	require.NoError(err)
	defer ethClient.Close()

	blockNumber, err := ethClient.BlockNumber(ctx)
	require.NoError(err)
	require.Equal(blockNumber, uint64(0))

	grpcConn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(err)
	defer grpcConn.Close()

	grpcShardClient := pb.NewWriteShardClient(grpcConn)

	for i := 1; i <= 20; i++ {
		blockFile, err := blockDataFiles.Open(fmt.Sprintf("blockdata/%d.bin", i))
		require.NoError(err)
		defer blockFile.Close()

		blockBytes, err := io.ReadAll(blockFile)
		require.NoError(err)

		_, err = grpcShardClient.ExecuteBlock(ctx, &pb.ExecuteBlockRequest{
			BlockBytes: blockBytes,
		})
		require.NoError(err)

		blockNumber, err = ethClient.BlockNumber(ctx)
		require.NoError(err)
		require.Equal(blockNumber, uint64(i))
	}
}
