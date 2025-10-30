// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockdb

import (
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/leveldb"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/prometheus/client_golang/prometheus"
)

const maxHeightKey = "max_height"

var maxHeightKeyBytes = []byte(maxHeightKey)

type BlockDB struct {
	lock      sync.RWMutex
	db        database.Database
	maxHeight uint64
}

type options struct {
	setMaxHeight *uint64
}

type Option func(*options)

func WithMaxHeight(maxHeight uint64) Option {
	return func(opts *options) {
		opts.setMaxHeight = &maxHeight
	}
}

func NewBlockDB(dbDir string, opts ...Option) (*BlockDB, error) {
	var options options
	for _, opt := range opts {
		opt(&options)
	}
	db, err := leveldb.New(dbDir, nil, logging.NoLog{}, prometheus.NewRegistry())
	if err != nil {
		return nil, fmt.Errorf("failed to create leveldb block database from %q: %w", dbDir, err)
	}

	b := &BlockDB{db: db}
	if err := b.setInitialMaxHeight(&options); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *BlockDB) setInitialMaxHeight(options *options) error {
	if options.setMaxHeight != nil {
		b.maxHeight = *options.setMaxHeight
		return nil
	}

	maxHeightBytes, err := b.db.Get(maxHeightKeyBytes)
	if err != nil && err != database.ErrNotFound {
		return fmt.Errorf("failed to read max height from block database: %w", err)
	}
	if err == database.ErrNotFound {
		b.maxHeight = 0
	} else {
		b.maxHeight = binary.BigEndian.Uint64(maxHeightBytes)
	}

	return nil

}

func (b *BlockDB) WriteBlock(height uint64, bytes []byte) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	batch := b.db.NewBatch()

	if height > b.maxHeight {
		b.maxHeight = height
		if err := batch.Put(maxHeightKeyBytes, blockKey(b.maxHeight)); err != nil {
			return fmt.Errorf("failed to update max height to %d: %w", b.maxHeight, err)
		}
	}

	if err := batch.Put(blockKey(height), bytes); err != nil {
		return fmt.Errorf("failed to put block at height %d: %w", height, err)
	}
	if err := batch.Write(); err != nil {
		return fmt.Errorf("failed to write block batch at height %d: %w", height, err)
	}
	return nil
}

func (b *BlockDB) GetBlockByHeight(height uint64) ([]byte, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	bytes, err := b.db.Get(blockKey(height))
	if err != nil {
		return nil, fmt.Errorf("failed to read block at height %d: %w", height, err)
	}
	return bytes, nil
}

func (b *BlockDB) GetMaxHeight() uint64 {
	b.lock.RLock()
	defer b.lock.RUnlock()

	return b.maxHeight
}

func (b *BlockDB) NewIteratorFromHeight(height uint64) database.Iterator {
	return b.db.NewIteratorWithStartAndPrefix(blockKey(height), nil)
}

func (b *BlockDB) Close() error {
	return b.db.Close()
}

func blockKey(height uint64) []byte {
	return binary.BigEndian.AppendUint64(nil, height)
}
