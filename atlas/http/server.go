// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	DefaultShutdownTimeout = 10 * time.Second
	DefaultPort            = 0
)

// serveOptions holds configuration for ServeWithContext
type serveOptions struct {
	logger          logging.Logger
	port            int
	shutdownTimeout time.Duration
}

// ServeOption configures ServeWithContext
type ServeOption func(*serveOptions)

// WithLogger sets the logger for the server. Defaults to logging.NoLog{}.
func WithLogger(log logging.Logger) ServeOption {
	return func(o *serveOptions) {
		o.logger = log
	}
}

// WithPort sets the port for the server. Defaults to 0 (random available port).
func WithPort(port int) ServeOption {
	return func(o *serveOptions) {
		o.port = port
	}
}

// WithShutdownTimeout sets the timeout for graceful shutdown. Defaults to 10 seconds.
func WithShutdownTimeout(timeout time.Duration) ServeOption {
	return func(o *serveOptions) {
		o.shutdownTimeout = timeout
	}
}

func ServeWithContext(ctx context.Context, handler http.Handler, opts ...ServeOption) error {
	// Apply default options
	options := &serveOptions{
		logger:          logging.NoLog{},
		port:            DefaultPort,
		shutdownTimeout: DefaultShutdownTimeout,
	}

	// Apply provided options
	for _, opt := range opts {
		opt(options)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", options.port),
		Handler: handler,
	}
	shutdownDone := make(chan error)

	go func() {
		defer close(shutdownDone)
		<-ctx.Done()
		// Attempt graceful shutdown
		options.logger.Info("Shutting down HTTP server...")

		// Shutdown the HTTP server, allowing active requests to complete
		shutdownCtx, cancel := context.WithTimeout(context.Background(), options.shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			shutdownDone <- fmt.Errorf("failed to shutdown HTTP server gracefully: %w", err)
		}
		options.logger.Info("Finished shutting down HTTP server")
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed: %w", err)
	}

	return <-shutdownDone
}
