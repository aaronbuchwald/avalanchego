// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package shard

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/gorilla/mux"
)

type Server struct {
	shard  ReadShard
	router *mux.Router
}

func NewServer(ctx context.Context, shard ReadShard) (*Server, error) {
	s := &Server{
		shard:  shard,
		router: mux.NewRouter(),
	}
	if err := s.setupRoutes(ctx); err != nil {
		return nil, fmt.Errorf("failed to setup routes: %w", err)
	}
	return s, nil
}

func (s *Server) setupRoutes(ctx context.Context) error {
	handlers, err := s.shard.CreateHandlers(ctx)
	if err != nil {
		return fmt.Errorf("failed to create handlers for shard: %w", err)
	}

	for path, handler := range handlers {
		s.router.Handle(path, handler)
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func ServeShard(ctx context.Context, log logging.Logger, port int, shard ReadShard) error {
	shardServer, err := NewServer(ctx, shard)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: shardServer,
	}

	shutdownDone := make(chan error)

	go func() {
		<-ctx.Done()
		// Attempt graceful shutdown
		log.Info("Shutting down HTTP server...")

		// Shutdown the HTTP server, allowing active requests to complete
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			shutdownDone <- fmt.Errorf("failed to shutdown HTTP server gracefully: %w", err)
		}
		log.Info("Finished shutting down shard HTTP server")

		close(shutdownDone)
	}()

	// Start the HTTP server
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed: %w", err)
	}

	// Wait for shutdown to complete if triggered
	return <-shutdownDone
}
