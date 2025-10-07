// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/avalanchego/atlas/evm"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/atlas/vm"
	"github.com/ava-labs/avalanchego/tests"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

const (
	stateDirFlag   = "state-dir"
	portFlag       = "port"
	startBlockFlag = "start-block"
	endBlockFlag   = "end-block"
)

// serveShardCmd represents the serveShard command
var serveShardCmd = &cobra.Command{
	Use:   "serve-shard",
	Short: "Serve a shard",
	RunE:  runServeShard,
}

func init() {
	rootCmd.AddCommand(serveShardCmd)

	serveShardCmd.PersistentFlags().String(stateDirFlag, "", "The directory to store the state of the shard")
	serveShardCmd.PersistentFlags().Int(portFlag, 0, "The port to serve the shard on")
	serveShardCmd.PersistentFlags().Uint64(startBlockFlag, 0, "The start block of the shard")
	serveShardCmd.PersistentFlags().Uint64(endBlockFlag, math.MaxUint64, "The end block of the shard")
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
	startBlock, err := cmd.PersistentFlags().GetUint64(startBlockFlag)
	if err != nil {
		return fmt.Errorf("failed to get start block: %w", err)
	}
	endBlock, err := cmd.PersistentFlags().GetUint64(endBlockFlag)
	if err != nil {
		return fmt.Errorf("failed to get end block: %w", err)
	}
	log := tests.NewDefaultLogger("evm-shard")
	cChainVM, cleanup, err := evm.New(context.Background(), log, stateDir)
	if err != nil {
		return fmt.Errorf("failed to create VM: %w", err)
	}
	defer cleanup()

	vmShard := vm.NewVMShardAdapter(cChainVM, startBlock, endBlock)
	shardServer, err := shard.NewServer(vmShard)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	// create the http server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: shardServer,
	}

	// Channel to listen for interrupt signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Channel to signal server shutdown is complete
	shutdownDone := make(chan struct{})

	go func() {
		<-sigCh
		// Attempt graceful shutdown
		log.Info("interrupt signal received, shutting down VM and HTTP server")
		// Shutdown the HTTP server, allowing active requests to complete
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Error("failed to shutdown HTTP server gracefully", zap.Error(err))
		}

		if err := cChainVM.Shutdown(context.Background()); err != nil {
			log.Error("failed to shutdown VM", zap.Error(err))
		}

		close(shutdownDone)
	}()

	// Start the HTTP server
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed: %w", err)
	}

	// Wait for shutdown to complete if triggered
	<-shutdownDone
	return nil
}
