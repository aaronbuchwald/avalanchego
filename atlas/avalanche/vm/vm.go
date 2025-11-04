// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/leveldb"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/enginetest"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/validators/validatorstest"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms"
	"github.com/ava-labs/avalanchego/vms/metervm"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	statesyncclient "github.com/ava-labs/coreth/sync/client"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	_ block.ChainVM = (*AtlasVM)(nil)
	_ shard.Shard   = (*AtlasVM)(nil)
)

type VMParams struct {
	Name          string
	Factory       vms.Factory
	NetworkConfig NetworkConfig
	ConfigBytes   []byte
}

type NetworkConfig struct {
	NetworkID         uint32
	SubnetID          ids.ID
	ChainID           ids.ID
	NetworkUpgrades   upgrade.Config
	XChainID          ids.ID
	CChainID          ids.ID
	AVAXAssetID       ids.ID
	GenesisBytes      []byte
	UpgradeBytes      []byte
	ChainIDToSubnetID map[ids.ID]ids.ID
}

func NewVMParams(
	vmName string,
	factory vms.Factory,
	networkConfig NetworkConfig,
	configBytes []byte,
) VMParams {
	return VMParams{
		Name:          vmName,
		Factory:       factory,
		NetworkConfig: networkConfig,
		ConfigBytes:   configBytes,
	}
}

type AtlasVM struct {
	block.ChainVM

	params  VMParams
	sender  *enginetest.Sender
	snowCtx *snow.Context
	closeDB func() error
}

func NewAtlasVM(
	ctx context.Context,
	params VMParams,
	currentStateDir string,
) (*AtlasVM, error) {
	// Create the prefix gatherer passed to the VM and register it with the top-level,
	// labeled gatherer.
	// TODO: expose metrics server and optionally a collector in-process.
	prefixGatherer := metrics.NewPrefixGatherer()

	vmMultiGatherer := metrics.NewPrefixGatherer()
	if err := prefixGatherer.Register(fmt.Sprintf("avalanche_%s", params.Name), vmMultiGatherer); err != nil {
		return nil, fmt.Errorf("failed to register vmMultiGatherer: %w", err)
	}

	meterVMRegistry := prometheus.NewRegistry()
	if err := prefixGatherer.Register("avalanche_meterchainvm", meterVMRegistry); err != nil {
		return nil, fmt.Errorf("failed to register meterVMRegistry: %w", err)
	}

	// TODO: add back consensus metrics registry following the same pattern in vm_reexecute_test.go

	// Create VM from factory
	vmIntf, err := params.Factory.New(logging.NoLog{})
	if err != nil {
		return nil, fmt.Errorf("failed to create VM from factory: %w", err)
	}
	vm := vmIntf.(block.ChainVM)

	// Create BLS key for warp signing
	blsKey, err := localsigner.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create BLS key: %w", err)
	}

	blsPublicKey := blsKey.PublicKey()
	warpSigner := warp.NewSigner(blsKey, params.NetworkConfig.NetworkID, params.NetworkConfig.ChainID)

	// Create databases
	var (
		vmDBDir      = filepath.Join(currentStateDir, "db")
		chainDataDir = filepath.Join(currentStateDir, "chain-data-dir")
	)

	db, err := leveldb.New(vmDBDir, nil, logging.NoLog{}, prometheus.NewRegistry())
	if err != nil {
		return nil, fmt.Errorf("failed to create DB: %w", err)
	}

	sharedMemoryDB := prefixdb.New([]byte("sharedmemory"), db)
	atomicMemory := atomic.NewMemory(sharedMemoryDB)

	// Wrap VM with metervm
	vm = metervm.NewBlockVM(vm, meterVMRegistry)
	sender := &enginetest.Sender{}

	bcLookup := ids.NewAliaser()
	aliasErr := errors.Join(
		bcLookup.Alias(params.NetworkConfig.XChainID, "X"),
		bcLookup.Alias(params.NetworkConfig.CChainID, "C"),
		bcLookup.Alias(constants.PlatformChainID, "P"),
	)
	if aliasErr != nil {
		return nil, fmt.Errorf("failed to alias chains: %w", aliasErr)
	}

	snowCtx := &snow.Context{
		NetworkID:       params.NetworkConfig.NetworkID,
		SubnetID:        params.NetworkConfig.SubnetID,
		ChainID:         params.NetworkConfig.ChainID,
		NodeID:          ids.GenerateTestNodeID(),
		PublicKey:       blsPublicKey,
		NetworkUpgrades: params.NetworkConfig.NetworkUpgrades,

		XChainID:    params.NetworkConfig.XChainID,
		CChainID:    params.NetworkConfig.CChainID,
		AVAXAssetID: params.NetworkConfig.AVAXAssetID,

		Log:          tests.NewDefaultLogger("vm"),
		SharedMemory: atomicMemory.NewSharedMemory(params.NetworkConfig.ChainID),
		BCLookup:     bcLookup,
		Metrics:      vmMultiGatherer,

		WarpSigner: warpSigner,

		ValidatorState: &validatorstest.State{
			GetSubnetIDF: func(_ context.Context, chainID ids.ID) (ids.ID, error) {
				subnetID, ok := params.NetworkConfig.ChainIDToSubnetID[chainID]
				if ok {
					return subnetID, nil
				}
				return ids.Empty, fmt.Errorf("unknown chainID: %s", chainID)
			},
		},
		ChainDataDir: chainDataDir,
	}
	if err := vm.Initialize(
		ctx,
		snowCtx,
		prefixdb.New([]byte("vm"), db),
		params.NetworkConfig.GenesisBytes,
		params.NetworkConfig.UpgradeBytes,
		params.ConfigBytes,
		nil,
		sender,
	); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize VM: %w", err)
	}

	return &AtlasVM{
		ChainVM: vm,
		params:  params,
		snowCtx: snowCtx,
		sender:  sender,
		closeDB: db.Close,
	}, nil
}

func (v *AtlasVM) ExecuteBlock(ctx context.Context, blockBytes []byte) error {
	blk, err := v.ParseBlock(ctx, blockBytes)
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

// Shutdown overrides the ChainVM Shutdown function to close the underlying VM DB
// after shutting down the ChainVM.
func (v *AtlasVM) Shutdown(ctx context.Context) error {
	return errors.Join(
		v.ChainVM.Shutdown(ctx),
		func() error {
			if v.closeDB == nil {
				return nil
			}
			return v.closeDB()
		}(),
	)
}

// SplitAtHeight creates a VM based on one state directory and creates a new one at height N
// based off of that state directory.
func (v *AtlasVM) SplitAtHeight(
	ctx context.Context,
	targetHeight uint64,
	targetStateDir string,
) error {
	sourceVM := v

	targetVM, err := NewAtlasVM(ctx, v.params, targetStateDir)
	if err != nil {
		return fmt.Errorf("failed to create target VM: %w", err)
	}
	defer targetVM.Shutdown(ctx)

	var (
		sourceSender = sourceVM.sender
		sourceNodeID = sourceVM.snowCtx.NodeID
		targetSender = targetVM.sender
		targetNodeID = targetVM.snowCtx.NodeID
	)

	sourceSender.SendAppResponseF = func(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
		go targetVM.AppResponse(ctx, sourceNodeID, requestID, response)
		return nil
	}
	targetSender.SendAppRequestF = func(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, request []byte) error {
		return sourceVM.AppRequest(ctx, sourceNodeID, requestID, time.Now().Add(1*time.Second), request)
	}

	if err := targetVM.Connected(ctx, sourceNodeID, statesyncclient.StateSyncVersion); err != nil {
		return fmt.Errorf("failed to connect target VM to source VM: %w", err)
	}
	if err := sourceVM.Connected(ctx, targetNodeID, statesyncclient.StateSyncVersion); err != nil {
		return fmt.Errorf("failed to connect source VM to target VM: %w", err)
	}

	sourceSyncableVM := sourceVM.ChainVM.(block.StateSyncableVM)
	targetSyncableVM := targetVM.ChainVM.(block.StateSyncableVM)

	sourceStateSummary, err := sourceSyncableVM.GetStateSummary(ctx, targetHeight)
	if err != nil {
		return fmt.Errorf("failed to get state summary from source VM: %w", err)
	}

	targetStateSummary, err := targetSyncableVM.ParseStateSummary(ctx, sourceStateSummary.Bytes())
	if err != nil {
		return fmt.Errorf("failed to parse state summary from target VM: %w", err)
	}

	syncMode, err := targetStateSummary.Accept(ctx)
	if err != nil {
		return fmt.Errorf("failed to accept state summary from target VM: %w", err)
	}
	if syncMode != block.StateSyncStatic {
		return fmt.Errorf("expected state sync mode to be %s, got %s", block.StateSyncStatic, syncMode)
	}

	msg, err := targetVM.WaitForEvent(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for state sync done: %w", err)
	}
	if msg != common.StateSyncDone {
		return fmt.Errorf("expected state sync done message, got %s", msg)
	}

	// SetState should report any state sync client errors that occurred async and would not
	// have been reported via WaitForEvent as this callback is the first sync opportunity for the client
	// VM to report a fatal error.
	if err := targetVM.SetState(ctx, snow.Bootstrapping); err != nil {
		return fmt.Errorf("failed to set target VM to bootstrapping: %w", err)
	}
	return nil
}

// TODO: must switch from GetBlockIDAtHeight to a VM-specific method that accounts for state sync behavior. A state sync'd EVM
// node will have the 256 blocks prior to the checkpoint it sync'd, but will not have the required state.
// This should be pushed down to the EVM implementation and called from here or all collapsed into an EVM specific implementation
// and ignore other cases.
func (v *AtlasVM) Configure(ctx context.Context, heightRange shard.HeightRange, blockClient shard.BlockClient) error {
	_, err := v.ChainVM.GetBlockIDAtHeight(ctx, heightRange.Start)
	if err != nil {
		return fmt.Errorf("failed to get block ID at height range start %d: %w", heightRange.Start, err)
	}

	lastAcceptedID, err := v.ChainVM.LastAccepted(ctx)
	if err != nil {
		return fmt.Errorf("failed to get last accepted block ID: %w", err)
	}
	lastAcceptedBlock, err := v.ChainVM.GetBlock(ctx, lastAcceptedID)
	if err != nil {
		return fmt.Errorf("failed to get last accepted block: %w", err)
	}
	lastAcceptedHeight := lastAcceptedBlock.Height()

	// If the end marker is set, confirm the end block is available. If not, attempt to fill in using the
	// block client.
	if heightRange.End != nil {
		endBlock := *heightRange.End
		_, err := v.ChainVM.GetBlockIDAtHeight(ctx, endBlock)
		// If we have the end block, we have the full static range and have nothing to do.
		if err == nil {
			return nil
		}
		// If there's an unexpected error, return it.
		if err != nil && err != database.ErrNotFound {
			return fmt.Errorf("failed to get block ID at height range end %d: %w", endBlock, err)
		}
		// Otherwise, we must have hit database.ErrNotFound, and we must attempt to fill the static range
		// from the block client.
		maxBlockHeight, err := blockClient.GetMaxHeight(ctx)
		if err != nil {
			return fmt.Errorf("failed to get max block height: %w", err)
		}
		if maxBlockHeight < endBlock {
			return fmt.Errorf("max block height %d is less than desired end height %d", maxBlockHeight, endBlock)
		}

		// The block client appears to have the full static range, so we roll forward up to and including the end block.
		for height := lastAcceptedHeight + 1; height <= endBlock; height++ {
			block, err := blockClient.GetBlockByHeight(ctx, height)
			if err != nil {
				return fmt.Errorf("failed to get block by height %d: %w", height, err)
			}
			if err := v.ExecuteBlock(ctx, block); err != nil {
				return fmt.Errorf("failed to execute block %d: %w", height, err)
			}
		}
		return nil
	}

	// If the end marker is nil, then sync from last accepted block to tip continuously.
	blockResults, err := shard.CreateBlockStreamFromClient(ctx, blockClient, lastAcceptedHeight)
	if err != nil {
		return fmt.Errorf("failed to create block stream: %w", err)
	}
	return shard.IngestBlockStream(ctx, shard.NewShardBlockResultHandler(v), blockResults)
}
