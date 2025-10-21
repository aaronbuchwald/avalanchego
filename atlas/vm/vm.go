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
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
	"github.com/ava-labs/avalanchego/vms"
	"github.com/ava-labs/avalanchego/vms/metervm"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	_ block.ChainVM = (*AtlasVM)(nil)
	_ shard.Shard   = (*AtlasVM)(nil)
)

type VMParams struct {
	Factory         vms.Factory
	CurrentStateDir string
	VMMultiGatherer metrics.MultiGatherer
	MeterVMRegistry prometheus.Registerer
	NetworkConfig   NetworkConfig
	ConfigBytes     []byte
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
	currentStateDir string,
	networkConfig NetworkConfig,
	configBytes []byte,
) (*VMParams, error) {
	// Create the prefix gatherer passed to the VM and register it with the top-level,
	// labeled gatherer.
	prefixGatherer := metrics.NewPrefixGatherer()

	vmMultiGatherer := metrics.NewPrefixGatherer()
	if err := prefixGatherer.Register(fmt.Sprintf("avalanche_%s", vmName), vmMultiGatherer); err != nil {
		return nil, fmt.Errorf("failed to register vmMultiGatherer: %w", err)
	}

	meterVMRegistry := prometheus.NewRegistry()
	if err := prefixGatherer.Register("avalanche_meterchainvm", meterVMRegistry); err != nil {
		return nil, fmt.Errorf("failed to register meterVMRegistry: %w", err)
	}
	// TODO: add back consensus metrics

	params := &VMParams{
		Factory:         factory,
		VMMultiGatherer: vmMultiGatherer,
		MeterVMRegistry: meterVMRegistry,
		CurrentStateDir: currentStateDir,
		NetworkConfig:   networkConfig,
		ConfigBytes:     configBytes,
	}
	return params, nil
}

type AtlasVM struct {
	block.ChainVM

	sender  *enginetest.Sender
	snowCtx *snow.Context
	closeDB func() error
}

func NewAtlasVM(
	ctx context.Context,
	params *VMParams,
) (*AtlasVM, error) {
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
		vmDBDir      = filepath.Join(params.CurrentStateDir, "db")
		chainDataDir = filepath.Join(params.CurrentStateDir, "chain-data-dir")
	)

	db, err := leveldb.New(vmDBDir, nil, logging.NoLog{}, prometheus.NewRegistry())
	if err != nil {
		return nil, fmt.Errorf("failed to create DB: %w", err)
	}

	sharedMemoryDB := prefixdb.New([]byte("sharedmemory"), db)
	atomicMemory := atomic.NewMemory(sharedMemoryDB)

	// Wrap VM with metervm
	vm = metervm.NewBlockVM(vm, params.MeterVMRegistry)
	sender := &enginetest.Sender{}

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
		BCLookup:     ids.NewAliaser(),
		Metrics:      params.VMMultiGatherer,

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

// Split creates a VM based on one state directory and creates a new one at height N
// based off of that state directory.
func (v *AtlasVM) Split(
	ctx context.Context,
	targetVMParams *VMParams,
	targetHeight uint64,
) error {
	sourceVM := v

	targetVM, err := NewAtlasVM(ctx, targetVMParams)
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
		if nodeID != targetNodeID {
			return fmt.Errorf("attempted to send response to unexpected nodeID: %s", nodeID)
		}

		go targetVM.AppResponse(ctx, nodeID, requestID, response)
		return nil
	}
	targetSender.SendAppRequestF = func(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, request []byte) error {
		if nodeIDs.Len() != 1 && nodeIDs.Contains(sourceNodeID) {
			return fmt.Errorf("attempted to send request to unexpected set of nodeIDs: %v", nodeIDs)
		}

		return sourceVM.AppRequest(ctx, sourceNodeID, requestID, time.Now().Add(1*time.Second), request)
	}

	if err := targetVM.Connected(ctx, sourceNodeID, &version.Application{}); err != nil {
		return fmt.Errorf("failed to connect target VM to source VM: %w", err)
	}
	if err := sourceVM.Connected(ctx, targetNodeID, &version.Application{}); err != nil {
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
