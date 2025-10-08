// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ava-labs/avalanchego/atlas/shard"

	"github.com/spf13/cobra"
)

const (
	stateDirFlag = "state-dir"
	portFlag     = "port"
	logLevelFlag = "log-level"
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
}

func runServeShard(cmd *cobra.Command, args []string) error {
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

	readShard, err := shardFactory.New(ctx, log, stateDir)
	if err != nil {
		return fmt.Errorf("failed to create VM: %w", err)
	}
	defer readShard.Shutdown(ctx)

	return shard.ServeShard(ctx, log, port, readShard)
}

// contextWithDefaultSignals returns a context and cancellation function where the context is cancelled
// when SIGINT or SIGTERM are received.
func contextWithDefaultSignals(ctx context.Context) (context.Context, context.CancelFunc) {
	return contextWithSignals(ctx, syscall.SIGINT, syscall.SIGTERM)
}

// contextWithSignals returns a context and cancellation function where the context is cancelled
// when any of the provided signals are received.
func contextWithSignals(ctx context.Context, sig ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)

	go func() {
		defer cancel()
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			return
		}
	}()

	return ctx, cancel
}
