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
	stateDirFlag = "state-dir"
	portFlag     = "port"
	logLevelFlag = "log-level"
	grpcPortFlag = "grpc-port"
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
	cmd.PersistentFlags().Int(grpcPortFlag, 0, "The port to serve the gRPC server on")
}

func getServeShardFlags(cmd *cobra.Command) (stateDir string, port int, err error) {
	stateDir, err = cmd.PersistentFlags().GetString(stateDirFlag)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get state directory: %w", err)
	}
	port, err = cmd.PersistentFlags().GetInt(portFlag)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get port: %w", err)
	}
	return stateDir, port, nil
}

func runServeShard(cmd *cobra.Command, args []string) error {
	stateDir, port, err := getServeShardFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to get serve shard flags: %w", err)
	}
	if err := initLogger(cmd); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	var (
		isGRPCPortSet = cmd.PersistentFlags().Changed(grpcPortFlag)
		grpcPort      int
	)
	if isGRPCPortSet {
		grpcPort, err = cmd.PersistentFlags().GetInt(grpcPortFlag)
		if err != nil {
			return fmt.Errorf("failed to get gRPC port: %w", err)
		}
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

	eg := errgroup.Group{}

	eg.Go(func() error {
		return atlashttp.ServeWithContext(
			ctx,
			shardServer,
			atlashttp.WithLogger(log),
			atlashttp.WithPort(port),
		)
	})

	if isGRPCPortSet {
		eg.Go(func() error {
			grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
			if err != nil {
				return fmt.Errorf("failed to listen on gRPC server port %d: %w", grpcPort, err)
			}
			defer grpcListener.Close()

			shard.ServeGRPCShard(ctx, grpcListener, readShard)
			return nil
		})
	}

	return eg.Wait()
}
