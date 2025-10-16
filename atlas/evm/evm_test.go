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

// readBlockData reads blocks in the range [1,20] from the blockdata directory and returns
// them in sequential order.
func readBlockData(tb testing.TB) [][]byte {
	require := require.New(tb)

	blocks := make([][]byte, 20)
	for i := 1; i <= 20; i++ {
		blockFile, err := blockDataFiles.Open(fmt.Sprintf("blockdata/%d.bin", i))
		require.NoError(err)
		defer blockFile.Close()
		blockBytes, err := io.ReadAll(blockFile)
		require.NoError(err)
		blocks[i-1] = blockBytes
	}

	return blocks
}

// executeBlocks executes all of the blocks on the shard sequentially.
func executeBlocks(tb testing.TB, ctx context.Context, shard shard.WriteShard, blocks [][]byte) {
	require := require.New(tb)
	for _, blockBytes := range blocks {
		require.NoError(shard.ExecuteBlock(ctx, blockBytes))
	}
}

func setupWithMultipleShards(tb testing.TB, endBlocks []uint64) *shardTest {
	require := require.New(tb)
	ctx, cancel := context.WithCancel(context.Background())

	blocks := readBlockData(tb)
	shards := make([]*vmShard, len(endBlocks))
	for i, endBlock := range endBlocks {
		log := tests.NewDefaultLogger("test-evm-shard")
		var err error
		evmShard, err := New(ctx, log, tb.TempDir())
		require.NoError(err)
		tb.Cleanup(func() {
			cancel()
			require.NoError(evmShard.Shutdown(ctx))
		})
		shards[i] = evmShard

		executeBlocks(tb, ctx, evmShard, blocks[:endBlock])
	}

	return &shardTest{
		require: require,
		ctx:     ctx,
		cancel:  cancel,
		shards:  shards,
	}
}

func setup(tb testing.TB) *shardTest {
	return setupWithMultipleShards(tb, []uint64{0}) // Create a single shard at genesis
}

func TestReadShard(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
	defer cancel()

	shardServer, err := shard.NewServer(shardTest.ctx, vmShard)
	require.NoError(err)

	testServer := httptest.NewServer(shardServer)
	defer testServer.Close()

	client, err := ethclient.Dial(testServer.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	blockNumber, err := client.BlockNumber(ctx)
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

	client, err := ethclient.Dial(testServer.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	blockNumber, err := client.BlockNumber(ctx)
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

		blockNumber, err = client.BlockNumber(ctx)
		require.NoError(err)
		require.Equal(blockNumber, uint64(i))
	}
}

func TestReadShardsManual(t *testing.T) {
	shard0Tip, shard1Tip := uint64(10), uint64(20)
	tip := shard1Tip
	shardTest := setupWithMultipleShards(t, []uint64{shard0Tip, shard1Tip})
	require, ctx, cancel, shard0, shard1 := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0], shardTest.shards[1]
	defer cancel()

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
		// Confirm expected behavior for shard0 (include blocks <= shard0Tip, exclude blocks > shard0Tip)
		if i <= shard0Tip {
			block, err := shard0Client.BlockByNumber(ctx, big.NewInt(int64(i)))
			require.NoError(err)
			require.Equal(block.NumberU64(), i)
		} else {
			block, err := shard0Client.BlockByNumber(ctx, big.NewInt(int64(i)))
			require.ErrorContains(err, "cannot query unfinalized data")
			require.Nil(block)
		}
		// Confirm expected behavior for shard1 (include all blocks <= shard1Tip)
		if i <= shard1Tip {
			block, err := shard1Client.BlockByNumber(ctx, big.NewInt(int64(i)))
			require.NoError(err)
			require.Equal(block.NumberU64(), i)
		}
	}
}

func TestReadShardsWithRouter(t *testing.T) {
	shard0Tip, shard1Tip := uint64(10), uint64(20)
	tip := shard1Tip
	shardTest := setupWithMultipleShards(t, []uint64{shard0Tip, shard1Tip})
	require, ctx, cancel, shard0, shard1 := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0], shardTest.shards[1]
	defer cancel()

	shard0Server, err := shard.NewServer(shardTest.ctx, shard0)
	require.NoError(err)

	shard0TestServer := httptest.NewServer(shard0Server)
	defer shard0TestServer.Close()

	shard1Server, err := shard.NewServer(shardTest.ctx, shard1)
	require.NoError(err)

	shard1TestServer := httptest.NewServer(shard1Server)
	defer shard1TestServer.Close()

	router := NewRouter([]*APIShard{
		{
			Endpoint: shard0TestServer.URL + "/rpc",
			Start:    0,
			End:      shard0Tip,
		},
		{
			Endpoint: shard1TestServer.URL + "/rpc",
			Start:    shard0Tip,
			End:      shard1Tip,
		},
	})

	routerTestServer := httptest.NewServer(router)
	defer routerTestServer.Close()

	client, err := ethclient.Dial(routerTestServer.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	for i := uint64(1); i < tip; i++ {
		block, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		require.NoError(err)
		require.Equal(block.NumberU64(), i)
	}
}
