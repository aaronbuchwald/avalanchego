// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"embed"
	"fmt"
	"io"
	"math/big"
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
	shards  []*vmShard
}

func setupWithMultipleShards(tb testing.TB, numShards int) *shardTest {
	require := require.New(tb)
	ctx, cancel := context.WithCancel(context.Background())

	shards := make([]*vmShard, numShards)
	for i := 0; i < numShards; i++ {
		log := tests.NewDefaultLogger("test-evm-shard")
		var err error
		evmShard, err := New(ctx, log, tb.TempDir())
		require.NoError(err)
		tb.Cleanup(func() {
			cancel()
			require.NoError(evmShard.Shutdown(ctx))
		})
		shards[i] = evmShard
	}

	return &shardTest{
		require: require,
		ctx:     ctx,
		cancel:  cancel,
		shards:  shards,
	}
}

func setup(tb testing.TB) *shardTest {
	return setupWithMultipleShards(tb, 1)
}

func TestReadShard(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
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
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
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

func TestReadShards(t *testing.T) {
	shardTest := setupWithMultipleShards(t, 2)
	require, ctx, cancel, shard0, shard1 := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0], shardTest.shards[1]
	defer cancel()

	tip := uint64(20)
	shard0Tip, shard1Tip := tip/2, tip

	for i := uint64(1); i < tip; i++ {
		blockFile, err := blockDataFiles.Open(fmt.Sprintf("blockdata/%d.bin", i))
		require.NoError(err)
		defer blockFile.Close()

		blockBytes, err := io.ReadAll(blockFile)
		require.NoError(err)

		if i < shard0Tip {
			require.NoError(shard0.ExecuteBlock(ctx, blockBytes))
		}
		if i < shard1Tip {
			require.NoError(shard1.ExecuteBlock(ctx, blockBytes))
		}
	}

	shard0Server, err := shard.NewServer(shardTest.ctx, shard0)
	require.NoError(err)

	shard0TestServer := httptest.NewServer(shard0Server)
	defer shard0TestServer.Close()

	shard1Server, err := shard.NewServer(shardTest.ctx, shard1)
	require.NoError(err)

	shard1TestServer := httptest.NewServer(shard1Server)
	defer shard1TestServer.Close()

	shard0Client, err := ethclient.Dial(shard0TestServer.URL + "/rpc")
	require.NoError(err)
	defer shard0Client.Close()

	shard1Client, err := ethclient.Dial(shard1TestServer.URL + "/rpc")
	require.NoError(err)
	defer shard1Client.Close()

	for i := uint64(1); i < tip; i++ {
		if i < shard0Tip {
			block, err := shard0Client.BlockByNumber(ctx, big.NewInt(int64(i)))
			require.NoError(err)
			require.Equal(block.NumberU64(), i)
		} else {
			block, err := shard0Client.BlockByNumber(ctx, big.NewInt(int64(i)))
			require.ErrorContains(err, "cannot query unfinalized data")
			require.Nil(block)
		}
		if i < shard1Tip {
			block, err := shard1Client.BlockByNumber(ctx, big.NewInt(int64(i)))
			require.NoError(err)
			require.Equal(block.NumberU64(), i)
		}
	}
}
