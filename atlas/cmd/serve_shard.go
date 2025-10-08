// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/utils/logging"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

const (
	stateDirFlag = "state-dir"
	portFlag     = "port"
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

	log := tests.NewDefaultLogger("evm-shard")
	ctx, cancel := contextWithSignals(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	return serveShard(ctx, log, stateDir, port)
}

// serveShard creates and runs a shard server until the context is cancelled
func serveShard(ctx context.Context, log logging.Logger, stateDir string, port int) error {
	readShard, err := shardFactory.New(ctx, log, stateDir)
	if err != nil {
		return fmt.Errorf("failed to create VM: %w", err)
	}
	defer readShard.Shutdown(ctx)

	shardServer, err := shard.NewServer(ctx, readShard)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	// create the http server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: shardServer,
	}

	// Channel to signal server shutdown is complete
	shutdownDone := make(chan struct{})

	go func() {
		<-ctx.Done()
		// Attempt graceful shutdown
		log.Info("Shutting down HTTP server and VM...")

		// Shutdown the HTTP server, allowing active requests to complete
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Error("failed to shutdown HTTP server gracefully", zap.Error(err))
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
