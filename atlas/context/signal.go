// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package context

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithDefaultSignals returns a context and cancellation function where the context is cancelled
// when SIGINT or SIGTERM are received.
func WithDefaultSignals(ctx context.Context) (context.Context, context.CancelFunc) {
	return WithSignals(ctx, syscall.SIGINT, syscall.SIGTERM)
}

// WithSignals returns a context and cancellation function where the context is cancelled
// when any of the provided signals are received.
func WithSignals(ctx context.Context, sig ...os.Signal) (context.Context, context.CancelFunc) {
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
