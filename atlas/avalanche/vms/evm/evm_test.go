// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"embed"
	"fmt"
	"io"
	"math/big"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/atlas/avalanche/vm"
	"github.com/ava-labs/avalanchego/atlas/blockdb"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

//go:embed blockdata/*
var blockDataFiles embed.FS

var blockBytes [][]byte

type shardTest struct {
	require *require.Assertions
	ctx     context.Context
	cancel  context.CancelFunc
	shards  []*vm.AtlasVM
}

// blockClientAdapter adapts a blockdb.Client to the shard.BlockClient interface
type blockClientAdapter struct {
	client *blockdb.Client
}

func (a *blockClientAdapter) GetBlockByHeight(ctx context.Context, height uint64) ([]byte, error) {
	return a.client.GetBlockByHeight(ctx, height)
}

func (a *blockClientAdapter) GetMaxHeight(ctx context.Context) (uint64, error) {
	return a.client.GetMaxHeight(ctx)
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
	shards := make([]*vm.AtlasVM, len(endBlocks))
	for i, endBlock := range endBlocks {
		var err error
		evmShard, err := NewMainnetAtlasVM(ctx, tb.TempDir())
		require.NoError(err)
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

func assertStateAvailable(tb testing.TB, ctx context.Context, client *ethclient.Client, blockNumber uint64, stateAvailable bool) {
	require := require.New(tb)

	testAddr := common.HexToAddress("0x376c47978271565f56DEB45495afa69E59c16Ab2")
	_, err := client.NonceAt(ctx, testAddr, big.NewInt(int64(blockNumber)))
	if stateAvailable {
		require.NoError(err, "expected state to be available for block %d", blockNumber)
	} else {
		require.ErrorContains(err, "cannot query unfinalized data", "expected state to be unavailable for block %d", blockNumber)
	}
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

	assertStateAvailable(t, ctx, client, 0, true)
}

func TestActiveShard(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
	defer func() {
		cancel()
		require.NoError(vmShard.Shutdown(ctx))
	}()

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

	for i := 1; i <= 20; i++ {
		blockFile, err := blockDataFiles.Open(fmt.Sprintf("blockdata/%d.bin", i))
		require.NoError(err)
		defer blockFile.Close()

		blockBytes, err := io.ReadAll(blockFile)
		require.NoError(err)

		require.NoError(vmShard.ExecuteBlock(ctx, blockBytes))

		assertStateAvailable(t, ctx, client, uint64(i), true)
	}
}

func TestReadShardsManual(t *testing.T) {
	shard0Tip, shard1Tip := uint64(10), uint64(20)
	tip := shard1Tip
	shardTest := setupWithMultipleShards(t, []uint64{shard0Tip, shard1Tip})
	require, ctx, cancel, shard0, shard1 := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0], shardTest.shards[1]
	defer func() {
		cancel()
		require.NoError(shard0.Shutdown(ctx))
		require.NoError(shard1.Shutdown(ctx))
	}()

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
			assertStateAvailable(t, ctx, shard0Client, uint64(i), true)
		} else {
			assertStateAvailable(t, ctx, shard0Client, uint64(i), false)
		}
		// Confirm expected behavior for shard1 (include all blocks <= shard1Tip)
		if i <= shard1Tip {
			assertStateAvailable(t, ctx, shard1Client, uint64(i), true)
		}
	}
}

func TestReadShardsWithRouter(t *testing.T) {
	shard0Tip, shard1Tip := uint64(10), uint64(20)
	tip := shard1Tip
	shardTest := setupWithMultipleShards(t, []uint64{shard0Tip, shard1Tip})
	require, ctx, cancel, shard0, shard1 := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0], shardTest.shards[1]
	defer func() {
		cancel()
		require.NoError(shard0.Shutdown(ctx))
		require.NoError(shard1.Shutdown(ctx))
	}()

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
			HeightRange: shard.HeightRange{
				Start: 0,
				End:   &shard0Tip,
			},
		},
		{
			Endpoint: shard1TestServer.URL + "/rpc",
			HeightRange: shard.HeightRange{
				Start: shard0Tip,
				End:   &shard1Tip,
			},
		},
	})

	routerTestServer := httptest.NewServer(router)
	defer routerTestServer.Close()

	client, err := ethclient.Dial(routerTestServer.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	for i := uint64(1); i < tip; i++ {
		assertStateAvailable(t, ctx, client, uint64(i), true)
	}
}

func TestSplitAtHeight(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
	defer func() {
		cancel()
		require.NoError(vmShard.Shutdown(ctx))
	}()

	blocks := readBlockData(t)
	executeBlocks(t, ctx, vmShard, blocks[:10])

	targetStateDir := t.TempDir()
	require.NoError(vmShard.SplitAtHeight(ctx, 10, targetStateDir))

	freshVMParams, err := newCChainArchiveVMParams(constants.MainnetID)
	require.NoError(err)

	targetVM, err := vm.NewAtlasVM(ctx, freshVMParams, targetStateDir)
	require.NoError(err)
	defer targetVM.Shutdown(ctx)

	executeBlocks(t, ctx, targetVM, blocks[10:])

	targetShardServer, err := shard.NewServer(ctx, targetVM)
	require.NoError(err)

	server := httptest.NewServer(targetShardServer)
	defer server.Close()

	client, err := ethclient.Dial(server.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	for i := uint64(10); i <= 20; i++ {
		assertStateAvailable(t, ctx, client, uint64(i), true)
	}
}

func TestConfigureStaticGenesisRange(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
	defer func() {
		cancel()
		require.NoError(vmShard.Shutdown(ctx))
	}()

	blocks := readBlockData(t)

	// Create a block database and populate it with blocks [1,10]
	blockDBInstance, err := blockdb.NewBlockDB(t.TempDir())
	require.NoError(err)
	defer blockDBInstance.Close()

	for i := uint64(1); i <= 10; i++ {
		require.NoError(blockDBInstance.WriteBlock(i, blocks[i-1]))
	}

	// Create a block client from the database
	blockDBServer := blockdb.NewHTTPHandler(blockDBInstance)
	testServer := httptest.NewServer(blockDBServer)
	defer testServer.Close()

	blockClient := &blockClientAdapter{client: blockdb.NewClient(testServer.URL)}

	// Configure the target VM with static range [0,10] (starting from genesis)
	startHeight := uint64(0)
	endHeight := uint64(10)
	heightRange := shard.HeightRange{
		Start: startHeight,
		End:   &endHeight,
	}

	require.NoError(vmShard.Configure(ctx, heightRange, blockClient))

	// Verify the VM has blocks [1,10]
	shardServer, err := shard.NewServer(ctx, vmShard)
	require.NoError(err)

	vmTestServer := httptest.NewServer(shardServer)
	defer vmTestServer.Close()

	client, err := ethclient.Dial(vmTestServer.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	for i := uint64(1); i <= 10; i++ {
		assertStateAvailable(t, ctx, client, i, true)
	}

	// Verify blocks beyond range are not available
	assertStateAvailable(t, ctx, client, 11, false)

	// Shutdown the VM to simulate a restart
	require.NoError(vmShard.Shutdown(ctx))

	// Re-create the VM and do not defer the shutdown, since it was already invoked on the same
	// instance above.
	vmShard, err = NewMainnetAtlasVM(ctx, t.TempDir())
	require.NoError(err)
	// Reconfigure with the same static range [0,10]
	require.NoError(vmShard.Configure(ctx, heightRange, blockClient))

	// Start a new server and a new client from the restarted VM
	restartedShardServer, err := shard.NewServer(ctx, vmShard)
	require.NoError(err)
	restartedVMTestServer := httptest.NewServer(restartedShardServer)
	defer restartedVMTestServer.Close()

	restartedClient, err := ethclient.Dial(restartedVMTestServer.URL + "/rpc")
	require.NoError(err)
	defer restartedClient.Close()

	// Confirm the static range [1,10] is still available
	for i := uint64(1); i <= 10; i++ {
		assertStateAvailable(t, ctx, restartedClient, i, true)
	}

	// Confirm block 11 is unavailable
	assertStateAvailable(t, ctx, restartedClient, 11, false)
}

func TestConfigureDynamicRange(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
	defer func() {
		cancel()
		require.NoError(vmShard.Shutdown(ctx))
	}()

	blocks := readBlockData(t)

	// Create a block database and populate it with blocks [1,10]
	blockDBInstance, err := blockdb.NewBlockDB(t.TempDir())
	require.NoError(err)
	defer blockDBInstance.Close()

	for i := uint64(1); i <= 10; i++ {
		require.NoError(blockDBInstance.WriteBlock(i, blocks[i-1]))
	}

	// Create a block client from the database
	blockDBServer := blockdb.NewHTTPHandler(blockDBInstance)
	testServer := httptest.NewServer(blockDBServer)
	defer testServer.Close()

	blockClient := &blockClientAdapter{client: blockdb.NewClient(testServer.URL)}

	// Configure the target VM with dynamic range [0, nil] (from block 1 to latest available)
	startHeight := uint64(0)
	heightRange := shard.HeightRange{
		Start: startHeight,
		End:   nil,
	}

	// XXX: improve testability of Configure as a dynamic range.
	configureCtx, configureCancel := context.WithTimeout(ctx, 3*time.Second)
	defer configureCancel()
	require.NoError(vmShard.Configure(configureCtx, heightRange, blockClient))

	// Verify the VM has blocks [1,10]
	shardServer, err := shard.NewServer(ctx, vmShard)
	require.NoError(err)

	vmTestServer := httptest.NewServer(shardServer)
	defer vmTestServer.Close()

	client, err := ethclient.Dial(vmTestServer.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	for i := uint64(1); i <= 10; i++ {
		assertStateAvailable(t, ctx, client, i, true)
	}

	// Verify block 11 is not available (only blocks [1,10] are in blockdb)
	assertStateAvailable(t, ctx, client, 11, false)
}

func TestConfigureStaticRangeAfterSplit(t *testing.T) {
	shardTest := setup(t)
	require, ctx, cancel, vmShard := shardTest.require, shardTest.ctx, shardTest.cancel, shardTest.shards[0]
	defer func() {
		cancel()
		require.NoError(vmShard.Shutdown(ctx))
	}()

	blocks := readBlockData(t)

	// Execute blocks [1,10] on the initial VM
	executeBlocks(t, ctx, vmShard, blocks[:10])

	// Split at height 10
	targetStateDir := t.TempDir()
	require.NoError(vmShard.SplitAtHeight(ctx, 10, targetStateDir))

	// Create a fresh VM from the split
	freshVMParams, err := newCChainArchiveVMParams(constants.MainnetID)
	require.NoError(err)

	targetVM, err := vm.NewAtlasVM(ctx, freshVMParams, targetStateDir)
	require.NoError(err)
	defer func() {
		require.NoError(targetVM.Shutdown(ctx))
	}()

	// Create a block database with blocks [11,20]
	blockDBInstance, err := blockdb.NewBlockDB(t.TempDir())
	require.NoError(err)
	defer blockDBInstance.Close()

	for i := uint64(11); i <= 20; i++ {
		require.NoError(blockDBInstance.WriteBlock(i, blocks[i-1]))
	}

	// Create a block client from the database
	blockDBServer := blockdb.NewHTTPHandler(blockDBInstance)
	testServer := httptest.NewServer(blockDBServer)
	defer testServer.Close()

	blockClient := &blockClientAdapter{client: blockdb.NewClient(testServer.URL)}

	// Configure the target VM with static range [10,20]
	// Note: Start should be where we split (10), end is 20
	startHeight := uint64(10)
	endHeight := uint64(20)
	heightRange := shard.HeightRange{
		Start: startHeight,
		End:   &endHeight,
	}

	require.NoError(targetVM.Configure(ctx, heightRange, blockClient))

	// Verify the VM has blocks [10,20]
	shardServer, err := shard.NewServer(ctx, targetVM)
	require.NoError(err)

	vmTestServer := httptest.NewServer(shardServer)
	defer vmTestServer.Close()

	client, err := ethclient.Dial(vmTestServer.URL + "/rpc")
	require.NoError(err)
	defer client.Close()

	for i := uint64(11); i <= 20; i++ {
		assertStateAvailable(t, ctx, client, i, true)
	}

	// // Shut down and restart the target VM to ensure persistence and block/state availability.
	require.NoError(targetVM.Shutdown(ctx))

	// Re-create targetVM in the same state directory (targetStateDir)
	targetVM, err = vm.NewAtlasVM(ctx, freshVMParams, targetStateDir)
	require.NoError(err)

	// // Reconfigure the height range and block client as before
	require.NoError(targetVM.Configure(ctx, heightRange, blockClient))

	// Start the server again for the restarted VM
	shardServer2, err := shard.NewServer(ctx, targetVM)
	require.NoError(err)

	vmTestServer2 := httptest.NewServer(shardServer2)
	defer vmTestServer2.Close()

	client2, err := ethclient.Dial(vmTestServer2.URL + "/rpc")
	require.NoError(err)
	defer client2.Close()

	for i := uint64(11); i <= 20; i++ {
		assertStateAvailable(t, ctx, client2, i, true)
	}
}
