// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/atlas/blockdb"
	atlascontext "github.com/ava-labs/avalanchego/atlas/context"
	atlashttp "github.com/ava-labs/avalanchego/atlas/http"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	stateDirFlag        = "state-dir"
	portFlag            = "port"
	logLevelFlag        = "log-level"
	startHeightFlag     = "start"
	endHeightFlag       = "end"
	blockServerAddrFlag = "block-server-addr"
)

// serveShardCmd represents the serveShard command
var serveShardCmd = &cobra.Command{
	Use:   "serve-shard",
	Short: "Serve a shard",
	RunE:  runServeShard,
}

func init() {
	rootCmd.AddCommand(serveShardCmd)

	serveShardCmd.PersistentFlags().String(logLevelFlag, "info", "The log level to use")
	registerServeShardFlags(serveShardCmd)
}

func registerServeShardFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(stateDirFlag, "", "The directory to store the state of the shard")
	cmd.PersistentFlags().Int(portFlag, 0, "The port to serve the shard on")
	cmd.PersistentFlags().Uint64(startHeightFlag, 0, "The start height to serve the shard from")
	cmd.PersistentFlags().Uint64(endHeightFlag, 0, "The end height to serve the shard to. If set to 0, the shard will attempt to serve blocks up to the latest available block.")
	cmd.PersistentFlags().String(blockServerAddrFlag, "", "The address of the block server to use")
}

func getServeShardFlags(cmd *cobra.Command) (stateDir string, port int, startHeight uint64, endHeight uint64, blockServerAddr string, err error) {
	stateDir, err = cmd.PersistentFlags().GetString(stateDirFlag)
	if err != nil {
		return "", 0, 0, 0, "", fmt.Errorf("failed to get state directory: %w", err)
	}
	port, err = cmd.PersistentFlags().GetInt(portFlag)
	if err != nil {
		return "", 0, 0, 0, "", fmt.Errorf("failed to get port: %w", err)
	}
	startHeight, err = cmd.PersistentFlags().GetUint64(startHeightFlag)
	if err != nil {
		return "", 0, 0, 0, "", fmt.Errorf("failed to get start height: %w", err)
	}
	endHeight, err = cmd.PersistentFlags().GetUint64(endHeightFlag)
	if err != nil {
		return "", 0, 0, 0, "", fmt.Errorf("failed to get end height: %w", err)
	}
	blockServerAddr, err = cmd.PersistentFlags().GetString(blockServerAddrFlag)
	if err != nil {
		return "", 0, 0, 0, "", fmt.Errorf("failed to get block server address: %w", err)
	}
	return stateDir, port, startHeight, endHeight, blockServerAddr, nil
}

func runServeShard(cmd *cobra.Command, args []string) error {
	stateDir, port, startHeight, endHeight, blockServerAddr, err := getServeShardFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to get serve shard flags: %w", err)
	}
	if err := initLogger(cmd); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	ctx, cancel := atlascontext.WithDefaultSignals(context.Background())
	defer cancel()

	readShard, err := shardFactory.New(ctx, log, stateDir)
	if err != nil {
		return fmt.Errorf("failed to create VM: %w", err)
	}
	defer readShard.Shutdown(ctx)

	shardServer, err := shard.NewServer(ctx, readShard)
	if err != nil {
		return fmt.Errorf("failed to create shard server: %w", err)
	}

	heightRange := shard.HeightRange{
		Start: startHeight,
	}
	if endHeight != 0 {
		heightRange.End = &endHeight
	}
	blockClient := blockdb.NewClient(blockServerAddr)
	if err != nil {
		return fmt.Errorf("failed to create block client: %w", err)
	}

	eg := errgroup.Group{}

	eg.Go(func() error {
		return atlashttp.ServeWithContext(
			ctx,
			shardServer,
			atlashttp.WithLogger(log),
			atlashttp.WithPort(port),
		)
	})
	eg.Go(func() error {
		return readShard.Configure(ctx, heightRange, blockClient)
	})

	return eg.Wait()
}
