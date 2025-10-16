// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/coreth/plugin/evm"
	"github.com/ava-labs/coreth/plugin/factory"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	mainnetXChainID    = ids.FromStringOrPanic("2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM")
	mainnetCChainID    = ids.FromStringOrPanic("2q9e4r6Mu3U68nU1fYjgbR6JvwrRx36CohpAX5UQxse55x1Q5")
	mainnetAvaxAssetID = ids.FromStringOrPanic("FvwEAhmxKfeiG8SnEvq42hc6whRyY3EFYAvebMqDNDGCgxN5Z")

	configBytes = []byte(`{
		"pruning-enabled": false
	}`)
)

func init() {
	evm.RegisterAllLibEVMExtras()
}

// createCChainMainnetVMParams creates VMParams with C-Chain mainnet configuration
func createCChainMainnetVMParams(
	log logging.Logger,
	currentStateDir string,
) (*VMParams, error) {
	// Create the prefix gatherer passed to the VM and register it with the top-level,
	// labeled gatherer.
	prefixGatherer := metrics.NewPrefixGatherer()

	vmMultiGatherer := metrics.NewPrefixGatherer()
	if err := prefixGatherer.Register("avalanche_evm", vmMultiGatherer); err != nil {
		return nil, fmt.Errorf("failed to register vmMultiGatherer: %w", err)
	}

	meterVMRegistry := prometheus.NewRegistry()
	if err := prefixGatherer.Register("avalanche_meterchainvm", meterVMRegistry); err != nil {
		return nil, fmt.Errorf("failed to register meterVMRegistry: %w", err)
	}

	// consensusRegistry includes the chain="C" label and the prefix "avalanche_snowman".
	// The consensus registry is passed to the executor to mimic a subset of consensus metrics.
	consensusRegistry := prometheus.NewRegistry()
	if err := prefixGatherer.Register("avalanche_snowman", consensusRegistry); err != nil {
		return nil, fmt.Errorf("failed to register consensusRegistry: %w", err)
	}

	// Get mainnet genesis configuration
	genesisConfig := genesis.GetConfig(constants.MainnetID)

	// Create chainIDToSubnetID mapping
	chainIDToSubnetID := map[ids.ID]ids.ID{
		mainnetXChainID: constants.PrimaryNetworkID,
		mainnetCChainID: constants.PrimaryNetworkID,
		ids.Empty:       constants.PrimaryNetworkID,
	}

	return &VMParams{
		Factory:           factory.Factory{},
		CurrentStateDir:   currentStateDir,
		VMMultiGatherer:   vmMultiGatherer,
		MeterVMRegistry:   meterVMRegistry,
		ChainIDToSubnetID: chainIDToSubnetID,
		NetworkID:         constants.MainnetID,
		SubnetID:          constants.PrimaryNetworkID,
		ChainID:           mainnetCChainID,
		NetworkUpgrades:   upgrade.Mainnet,
		XChainID:          mainnetXChainID,
		CChainID:          mainnetCChainID,
		AVAXAssetID:       mainnetAvaxAssetID,
		GenesisBytes:      []byte(genesisConfig.CChainGenesis),
		UpgradeBytes:      nil,
		ConfigBytes:       configBytes,
	}, nil
}
