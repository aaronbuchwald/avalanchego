// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/atlas/vm"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/coreth/plugin/evm"
	"github.com/ava-labs/coreth/plugin/factory"
)

var EVMShardFactory shard.ShardFactory = (*Factory)(nil)

var (
	mainnetXChainID    = ids.FromStringOrPanic("2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM")
	mainnetCChainID    = ids.FromStringOrPanic("2q9e4r6Mu3U68nU1fYjgbR6JvwrRx36CohpAX5UQxse55x1Q5")
	mainnetAvaxAssetID = ids.FromStringOrPanic("FvwEAhmxKfeiG8SnEvq42hc6whRyY3EFYAvebMqDNDGCgxN5Z")

	testnetCChainID    = ids.FromStringOrPanic("yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp")
	testnetXChainID    = ids.FromStringOrPanic("2JVSBoinj9C2J33VntvzYtVJNZdN2NKiwwKjcumHUWEb5DbBrm")
	testnetAVAXAssetID = ids.FromStringOrPanic("U8iRqJoiJm8xZHAacmvYyZVwqQx6uDNtQeP3CQ6fcgQk3JqnK")

	localCChainID    = ids.FromStringOrPanic("2owdGqyG6FFzTHy5qhenDXQcEghvr571KZE3gSfRJERSJinuwC")
	localXChainID    = ids.FromStringOrPanic("2eNy1mUFdmaxXNj1eQHUe7Np4gju9sJsEtWQ4MX3ToiNKuADed")
	localAVAXAssetID = ids.FromStringOrPanic("2fombhL7aGPwj3KH4bfrmJwW6PVnMobf9Y2fn9GwxiAAJyFDbe")

	mainnetGenesis = []byte(genesis.GetConfig(constants.MainnetID).CChainGenesis)
	testnetGenesis = []byte(genesis.GetConfig(constants.FujiID).CChainGenesis)
	localGenesis   = []byte(genesis.GetConfig(constants.LocalID).CChainGenesis)

	networkConfigMap = map[uint32]vm.NetworkConfig{
		constants.MainnetID: {
			NetworkID:       constants.MainnetID,
			SubnetID:        constants.PrimaryNetworkID,
			ChainID:         mainnetCChainID,
			NetworkUpgrades: upgrade.Mainnet,
			XChainID:        mainnetXChainID,
			CChainID:        mainnetCChainID,
			AVAXAssetID:     mainnetAvaxAssetID,
			ChainIDToSubnetID: map[ids.ID]ids.ID{
				mainnetXChainID: constants.PrimaryNetworkID,
				mainnetCChainID: constants.PrimaryNetworkID,
				ids.Empty:       constants.PrimaryNetworkID,
			},
			GenesisBytes: []byte(genesis.GetConfig(constants.MainnetID).CChainGenesis),
			UpgradeBytes: nil,
		},
		constants.FujiID: {
			NetworkID:       constants.FujiID,
			SubnetID:        constants.PrimaryNetworkID,
			ChainID:         testnetCChainID,
			NetworkUpgrades: upgrade.Fuji,
			XChainID:        testnetXChainID,
			CChainID:        testnetCChainID,
			AVAXAssetID:     testnetAVAXAssetID,
			ChainIDToSubnetID: map[ids.ID]ids.ID{
				testnetXChainID: constants.PrimaryNetworkID,
				ids.Empty:       constants.PrimaryNetworkID,
				testnetCChainID: constants.PrimaryNetworkID,
			},
			GenesisBytes: []byte(genesis.GetConfig(constants.FujiID).CChainGenesis),
			UpgradeBytes: nil,
		},
		constants.LocalID: {
			NetworkID:       constants.LocalID,
			SubnetID:        constants.PrimaryNetworkID,
			ChainID:         localCChainID,
			NetworkUpgrades: upgrade.Default,
			XChainID:        localXChainID,
			CChainID:        localCChainID,
			AVAXAssetID:     localAVAXAssetID,
			ChainIDToSubnetID: map[ids.ID]ids.ID{
				localXChainID: constants.PrimaryNetworkID,
				ids.Empty:     constants.PrimaryNetworkID,
				localCChainID: constants.PrimaryNetworkID,
			},
			GenesisBytes: []byte(genesis.GetConfig(constants.LocalID).CChainGenesis),
			UpgradeBytes: nil,
		},
	}

	configBytes = []byte(`{
		"pruning-enabled": false,
		"state-sync-enabled": true,
		"state-sync-commit-interval": 1,
		"commit-interval": 1,
		"state-sync-min-blocks": 1
	}`)
)

func init() {
	evm.RegisterAllLibEVMExtras()
}

func newCChainArchiveVMParams(
	networkID uint32,
	log logging.Logger,
	currentStateDir string,
) (*vm.VMParams, error) {
	networkConfig, ok := networkConfigMap[networkID]
	if !ok {
		return nil, fmt.Errorf("unknown networkID: %d", networkID)
	}
	return vm.NewVMParams(
		"evm",
		&factory.Factory{},
		currentStateDir,
		networkConfig,
		configBytes,
	)
}

type Factory struct{}

func (f *Factory) New(ctx context.Context, log logging.Logger, stateDir string) (shard.Shard, error) {
	return NewMainnetAtlasVM(ctx, log, stateDir)
}

func NewMainnetAtlasVM(
	ctx context.Context,
	log logging.Logger,
	stateDir string,
) (*vm.AtlasVM, error) {
	vmParams, err := newCChainArchiveVMParams(constants.MainnetID, log, stateDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM params: %w", err)
	}
	return vm.NewAtlasVM(ctx, vmParams)
}
