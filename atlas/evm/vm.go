// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database/leveldb"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/enginetest"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/validators/validatorstest"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms"
	"github.com/ava-labs/avalanchego/vms/metervm"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/prometheus/client_golang/prometheus"
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

// CreateVM creates a new VM instance from the provided VMParams
func CreateVM(
	ctx context.Context,
	params *VMParams,
) (block.ChainVM, func() error, error) {
	// Create VM from factory
	vmIntf, err := params.Factory.New(logging.NoLog{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create VM from factory: %w", err)
	}
	vm := vmIntf.(block.ChainVM)

	// Create BLS key for warp signing
	blsKey, err := localsigner.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create BLS key: %w", err)
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
		return nil, nil, fmt.Errorf("failed to create DB: %w", err)
	}

	sharedMemoryDB := prefixdb.New([]byte("sharedmemory"), db)
	atomicMemory := atomic.NewMemory(sharedMemoryDB)

	// Wrap VM with metervm
	vm = metervm.NewBlockVM(vm, params.MeterVMRegistry)

	if err := vm.Initialize(
		ctx,
		&snow.Context{
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
		},
		prefixdb.New([]byte("vm"), db),
		params.NetworkConfig.GenesisBytes,
		params.NetworkConfig.UpgradeBytes,
		params.ConfigBytes,
		nil,
		&enginetest.Sender{},
	); err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("failed to initialize VM: %w", err)
	}

	return vm, db.Close, nil
}
