// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package shard

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	shard  Shard
	router *mux.Router
}

func NewServer(shard Shard) (*Server, error) {
	s := &Server{
		shard:  shard,
		router: mux.NewRouter(),
	}
	if err := s.setupRoutes(context.Background()); err != nil {
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
