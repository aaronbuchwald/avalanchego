// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/ava-labs/avalanchego/atlas/shard"
)

const (
	grpcServerPortFlag = "grpc-server-port"
)

var activeShardCmd = &cobra.Command{
	Use:   "active-shard",
	Short: "Run an active shard listening for blocks to execute at tip",
	RunE:  runActiveShard,
}

func init() {
	rootCmd.AddCommand(activeShardCmd)

	activeShardCmd.PersistentFlags().String(logLevelFlag, "info", "The log level to use")
	registerServeShardFlags(activeShardCmd)
	registerActiveShardFlags(activeShardCmd)
}

func registerActiveShardFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Int(grpcServerPortFlag, 9000, "The port to serve the gRPC server on")
}

func runActiveShard(cmd *cobra.Command, args []string) error {
	stateDir, err := cmd.PersistentFlags().GetString(stateDirFlag)
	if err != nil {
		return fmt.Errorf("failed to get state directory: %w", err)
	}
	port, err := cmd.PersistentFlags().GetInt(portFlag)
	if err != nil {
		return fmt.Errorf("failed to get port: %w", err)
	}
	if err := initLogger(cmd); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	ctx, cancel := contextWithDefaultSignals(context.Background())
	defer cancel()

	activeShard, err := shardFactory.New(ctx, log, stateDir)
	if err != nil {
		return fmt.Errorf("failed to create VM: %w", err)
	}
	defer activeShard.Shutdown(ctx)

	grpcPort, err := cmd.PersistentFlags().GetInt(grpcServerPortFlag)
	if err != nil {
		return fmt.Errorf("failed to get gRPC server port: %w", err)
	}
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen on gRPC server port %d: %w", port, err)
	}

	eg := errgroup.Group{}
	eg.Go(func() error {
		shard.ServeGRPCShard(ctx, grpcListener, activeShard)
		return nil
	})
	eg.Go(func() error {
		shard.ServeShard(ctx, log, port, activeShard)
		return nil
	})
	return eg.Wait()
}
