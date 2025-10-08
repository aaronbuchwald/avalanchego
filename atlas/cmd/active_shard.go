// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	blockDirFlag         = "block-dir"
	blockStreamStartFlag = "block-stream-start"
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

func registerActiveShardFlags(cmd *cobra.Command) {}

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

	readShard, err := shardFactory.New(ctx, log, stateDir)
	if err != nil {
		return fmt.Errorf("failed to create VM: %w", err)
	}
	defer readShard.Shutdown(ctx)

	return serveShard(ctx, log, readShard, port)
}
