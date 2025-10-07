package evm

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/leveldb"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/enginetest"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/validators/validatorstest"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/metervm"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
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

type VMParams struct {
	vmAndSharedMemoryDB database.Database
	chainDataDir        string
	configBytes         []byte
	vmMultiGatherer     metrics.MultiGatherer
	meterVMRegistry     prometheus.Registerer
}

func NewVMParams(log logging.Logger, currentStateDir string, configBytes []byte) (*VMParams, func() error, error) {
	// Create the prefix gatherer passed to the VM and register it with the top-level,
	// labeled gatherer.
	prefixGatherer := metrics.NewPrefixGatherer()

	vmMultiGatherer := metrics.NewPrefixGatherer()
	if err := prefixGatherer.Register("avalanche_evm", vmMultiGatherer); err != nil {
		return nil, nil, fmt.Errorf("failed to register vmMultiGatherer: %w", err)
	}

	meterVMRegistry := prometheus.NewRegistry()
	if err := prefixGatherer.Register("avalanche_meterchainvm", meterVMRegistry); err != nil {
		return nil, nil, fmt.Errorf("failed to register meterVMRegistry: %w", err)
	}

	// consensusRegistry includes the chain="C" label and the prefix "avalanche_snowman".
	// The consensus registry is passed to the executor to mimic a subset of consensus metrics.
	consensusRegistry := prometheus.NewRegistry()
	if err := prefixGatherer.Register("avalanche_snowman", consensusRegistry); err != nil {
		return nil, nil, fmt.Errorf("failed to register consensusRegistry: %w", err)
	}

	// TODO: optionally collect metrics
	// if metricsEnabled {
	// 	collectRegistry(b, log, "c-chain-reexecution", prefixGatherer, labels)
	// }

	var (
		vmDBDir      = filepath.Join(currentStateDir, "db")
		chainDataDir = filepath.Join(currentStateDir, "chain-data-dir")
	)

	db, err := leveldb.New(vmDBDir, nil, log, prometheus.NewRegistry())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create DB: %w", err)
	}
	return &VMParams{
		vmAndSharedMemoryDB: db,
		chainDataDir:        chainDataDir,
		configBytes:         configBytes,
		vmMultiGatherer:     vmMultiGatherer,
		meterVMRegistry:     meterVMRegistry,
	}, db.Close, nil
}

func NewFromParams(
	ctx context.Context,
	params *VMParams,
) (block.ChainVM, error) {
	factory := factory.Factory{}
	vmIntf, err := factory.New(logging.NoLog{})
	if err != nil {
		return nil, fmt.Errorf("failed to create VM from factory: %w", err)
	}
	vm := vmIntf.(block.ChainVM)

	blsKey, err := localsigner.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create BLS key: %w", err)
	}

	blsPublicKey := blsKey.PublicKey()
	warpSigner := warp.NewSigner(blsKey, constants.MainnetID, mainnetCChainID)

	genesisConfig := genesis.GetConfig(constants.MainnetID)

	sharedMemoryDB := prefixdb.New([]byte("sharedmemory"), params.vmAndSharedMemoryDB)
	atomicMemory := atomic.NewMemory(sharedMemoryDB)

	chainIDToSubnetID := map[ids.ID]ids.ID{
		mainnetXChainID: constants.PrimaryNetworkID,
		mainnetCChainID: constants.PrimaryNetworkID,
		ids.Empty:       constants.PrimaryNetworkID,
	}

	vm = metervm.NewBlockVM(vm, params.meterVMRegistry)

	if err := vm.Initialize(
		ctx,
		&snow.Context{
			NetworkID:       constants.MainnetID,
			SubnetID:        constants.PrimaryNetworkID,
			ChainID:         mainnetCChainID,
			NodeID:          ids.GenerateTestNodeID(),
			PublicKey:       blsPublicKey,
			NetworkUpgrades: upgrade.Mainnet,

			XChainID:    mainnetXChainID,
			CChainID:    mainnetCChainID,
			AVAXAssetID: mainnetAvaxAssetID,

			Log:          tests.NewDefaultLogger("mainnet-vm-reexecution"),
			SharedMemory: atomicMemory.NewSharedMemory(mainnetCChainID),
			BCLookup:     ids.NewAliaser(),
			Metrics:      params.vmMultiGatherer,

			WarpSigner: warpSigner,

			ValidatorState: &validatorstest.State{
				GetSubnetIDF: func(_ context.Context, chainID ids.ID) (ids.ID, error) {
					subnetID, ok := chainIDToSubnetID[chainID]
					if ok {
						return subnetID, nil
					}
					return ids.Empty, fmt.Errorf("unknown chainID: %s", chainID)
				},
			},
			ChainDataDir: params.chainDataDir,
		},
		prefixdb.New([]byte("vm"), params.vmAndSharedMemoryDB),
		[]byte(genesisConfig.CChainGenesis),
		nil,
		params.configBytes,
		nil,
		&enginetest.Sender{},
	); err != nil {
		return nil, fmt.Errorf("failed to initialize VM: %w", err)
	}

	return vm, nil
}

func New(
	ctx context.Context,
	log logging.Logger,
	currentStateDir string,
) (block.ChainVM, func() error, error) {
	params, close, err := NewVMParams(log, currentStateDir, configBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create VM params: %w", err)
	}
	vm, err := NewFromParams(ctx, params)
	if err != nil {
		close()
		return nil, nil, fmt.Errorf("failed to create VM: %w", err)
	}
	return vm, close, nil
}
