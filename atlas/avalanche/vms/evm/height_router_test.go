// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMethodToHeightParamIndexLoaded(t *testing.T) {
	// Verify that the JSON was loaded correctly
	require.NotNil(t, methodToHeightParamIndex)
	require.Greater(t, len(methodToHeightParamIndex), 0, "methodToHeightParamIndex should not be empty")

	// Test some known mappings
	expectedMappings := map[string]int{
		"eth_getBlockByNumber":     0,
		"eth_getTransactionCount":  1,
		"eth_getBalance":           1,
		"eth_call":                 1,
		"debug_traceBlockByNumber": 0,
		"debug_traceCall":          1,
	}

	for method, expectedIndex := range expectedMappings {
		actualIndex, exists := methodToHeightParamIndex[method]
		require.True(t, exists, "method %s should exist in methodToHeightParamIndex", method)
		require.Equal(t, expectedIndex, actualIndex, "method %s should have index %d but got %d", method, expectedIndex, actualIndex)
	}

	// Verify we have a reasonable number of methods (at least 20)
	require.GreaterOrEqual(t, len(methodToHeightParamIndex), 20, "should have at least 20 methods mapped")
}
