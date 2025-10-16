// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/coreth/rpc"
)

var (
	_               http.Handler = (*Router)(nil)
	errNoShardFound              = errors.New("no shard found for height")
)

// TODO: make boundaries dynamic to support active + archival process w/o restart
type APIShard struct {
	Start    uint64 // Start block of the shard's range. Never changes.
	End      uint64 // End block of the shard's range. 0 indicates the shard is active and contains Start to tip.
	Endpoint string
}

// Atlas assumes that it only handles queries that access state at a specific height. All other queries can be handled without
// enabling archive mode.
// To that end, all such queries use their last parameter in the standard Ethereum JSON RPC calls to encode either a block hash
// or number, so we decode that parameter, perform the request on the required shard, and relay it to the client.
// Atlas could include a separate class internally to handle such requests and perhaps should to provide a complete Ethereum JSON RPC
// interface, but for now we handle the hard part (archival storage) and defer a complete RPC interface implementation.
type Router struct {
	log    logging.Logger
	shards []*APIShard
}

func NewRouter(shards []*APIShard) *Router {
	return &Router{
		log:    logging.NoLog{},
		shards: shards,
	}
}

// paramsOnlyRequest is a struct that contains only the params field of a JSON RPC request
// because this is the only field from the body required to extract the height.
type paramsOnlyRequest struct {
	Params []interface{} `json:"params"`
}

func extractHeightFromParams(params []interface{}) (uint64, error) {
	if len(params) == 0 {
		return 0, errors.New("no params found")
	}

	lastParam := params[len(params)-1]
	switch lastParam.(type) {
	case string:
		bnh := rpc.BlockNumberOrHash{}
		err := bnh.UnmarshalJSON([]byte(lastParam.(string)))
		if err != nil {
			return 0, err
		}
		blockNumber, ok := bnh.Number()
		if !ok {
			return 0, fmt.Errorf("invalid block number or hash %s", lastParam.(string))
		}
		blockNumberInt64 := blockNumber.Int64()
		if blockNumberInt64 < 0 {
			return 0, fmt.Errorf("special case block numbers not supported %s", lastParam.(string))
		}
		return uint64(blockNumberInt64), nil
	default:
		return 0, fmt.Errorf("invalid type for block hash or number: %T", params[len(params)-1])
	}
}

// findShard returns the first shard that contains height or nil of no such shard exists
func findShard(height uint64, shards []*APIShard) *APIShard {
	for _, shard := range shards {
		if height >= shard.Start && (height <= shard.End || shard.End == 0) {
			return shard
		}
	}
	return nil
}

// ServeHTTP implements http.Handler to forward requests to the correct shard based on height.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Read the full body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Restore the body for forwarding later
	req.Body = io.NopCloser(bytes.NewReader(body))

	// Unmarshal to extract params
	var paramsReq paramsOnlyRequest
	if err := json.Unmarshal(body, &paramsReq); err != nil {
		http.Error(w, "failed to unmarshal params: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Extract height
	height, err := extractHeightFromParams(paramsReq.Params)
	if err != nil {

		http.Error(w, "failed to extract height: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Find the correct shard
	shard := findShard(height, r.shards)
	if shard == nil {
		// TODO: incorrect auto-generated code, this needs to return an error message specific to the call
		// if the height is not found.
		http.Error(w, errNoShardFound.Error(), http.StatusNotFound)
		return
	}

	// Forward the request to the shard
	// Assume each apiShard has a ServeHTTP method or an http.Handler
	// If not, you may need to implement a forwarding mechanism (e.g., HTTP client to shard's endpoint)
	// Forward the request to the shard's endpoint using http.Client
	proxyReq, err := http.NewRequestWithContext(req.Context(), req.Method, shard.Endpoint, bytes.NewReader(body))
	if err != nil {
		http.Error(w, "failed to create proxy request: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Copy headers
	for k, vv := range req.Header {
		for _, v := range vv {
			proxyReq.Header.Add(k, v)
		}
	}
	// Use http.DefaultClient to send the request
	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, "failed to forward request to shard: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	// Copy response headers
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		panic(fmt.Errorf("failed to copy response body: %w", err))
	}
}
