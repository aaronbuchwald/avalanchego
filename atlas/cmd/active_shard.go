// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"
	"net"

	atlascontext "github.com/ava-labs/avalanchego/atlas/context"
	atlashttp "github.com/ava-labs/avalanchego/atlas/http"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
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

func getActiveShardFlags(cmd *cobra.Command) (grpcServerPort int, err error) {
	grpcServerPort, err = cmd.PersistentFlags().GetInt(grpcServerPortFlag)
	if err != nil {
		return 0, fmt.Errorf("failed to get gRPC server port: %w", err)
	}
	return grpcServerPort, nil
}

func runActiveShard(cmd *cobra.Command, args []string) error {
	stateDir, port, err := getServeShardFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to get state directory: %w", err)
	}
	if err := initLogger(cmd); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	grpcPort, err := getActiveShardFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to get gRPC server port: %w", err)
	}

	ctx, cancel := atlascontext.WithDefaultSignals(context.Background())
	defer cancel()

	activeShard, err := shardFactory.New(ctx, log, stateDir)
	if err != nil {
		return fmt.Errorf("failed to create VM: %w", err)
	}
	defer activeShard.Shutdown(ctx)

	shardServer, err := shard.NewServer(ctx, activeShard)
	if err != nil {
		return fmt.Errorf("failed to create shard server: %w", err)
	}

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen on gRPC server port %d: %w", grpcPort, err)
	}

	eg := errgroup.Group{}
	eg.Go(func() error {
		shard.ServeGRPCShard(ctx, grpcListener, activeShard)
		return nil
	})
	eg.Go(func() error {
		return atlashttp.ServeWithContext(
			ctx,
			shardServer,
			atlashttp.WithLogger(log),
			atlashttp.WithPort(port),
		)
	})
	return eg.Wait()
}
