// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockdb

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/memdb"
)

func newTestBlockDB() *BlockDB {
	db := memdb.New()
	return &BlockDB{db: db}
}

func TestServiceGetBlockByHeight(t *testing.T) {
	testCases := []struct {
		name          string
		height        uint64
		expectedBlock []byte
		expectedErr   error
	}{
		{
			name:          "genesis block",
			height:        0,
			expectedBlock: []byte("genesis block"),
		},
		{
			name:          "block 1",
			height:        1,
			expectedBlock: []byte("block 1 data"),
		},
		{
			name:          "block 100",
			height:        100,
			expectedBlock: []byte("block 100 with more complex data"),
		},
		{
			name:        "block not found",
			height:      101,
			expectedErr: database.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)

			blockDB := newTestBlockDB()

			if len(tc.expectedBlock) > 0 {
				require.NoError(blockDB.WriteBlock(tc.height, tc.expectedBlock))
			}

			handler := NewHTTPHandler(blockDB)
			server := httptest.NewServer(handler)
			defer server.Close()

			client := NewClient(server.URL)

			block, err := client.GetBlockByHeight(context.Background(), tc.height)
			if tc.expectedErr != nil {
				require.ErrorContains(err, tc.expectedErr.Error())
				require.Nil(block)
			} else {
				require.NoError(err)
				require.Equal(tc.expectedBlock, block)
			}
		})
	}
}
