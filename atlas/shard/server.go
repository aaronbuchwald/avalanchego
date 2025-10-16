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
	shard  ReadShard
	router *mux.Router
}

func NewServer(ctx context.Context, shard ReadShard) (*Server, error) {
	s := &Server{
		shard:  shard,
		router: mux.NewRouter(),
	}

	handlers, err := s.shard.CreateHandlers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create handlers for shard: %w", err)
	}
	for path, handler := range handlers {
		s.router.Handle(path, handler)
	}

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
