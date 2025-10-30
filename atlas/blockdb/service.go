// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockdb

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/ava-labs/avalanchego/utils/json"
	rpcClient "github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/gorilla/rpc/v2"
)

func NewHTTPHandler(db *BlockDB) http.Handler {
	server := rpc.NewServer()
	codec := json.NewCodec()
	server.RegisterCodec(codec, "application/json")
	server.RegisterCodec(codec, "application/json;charset=UTF-8")
	if err := server.RegisterService(&Blocks{db: db}, "blocks"); err != nil {
		panic(fmt.Sprintf("failed to register blocks service: %s", err))
	}
	return server
}

type Blocks struct{ db *BlockDB }

type GetBlockByHeightArgs struct {
	Height uint64 `json:"height"`
}

type GetBlockByHeightResponse struct {
	Block string `json:"block"`
}

func (s *Blocks) GetBlockByHeight(_ *http.Request, args *GetBlockByHeightArgs, reply *GetBlockByHeightResponse) error {
	block, err := s.db.GetBlockByHeight(args.Height)
	if err != nil {
		return err
	}

	reply.Block = hex.EncodeToString(block)
	return nil
}

type GetMaxHeightResponse struct {
	Height uint64 `json:"height"`
}

func (s *Blocks) GetMaxHeight(_ *http.Request, _ *struct{}, reply *GetMaxHeightResponse) error {
	reply.Height = s.db.GetMaxHeight()
	return nil
}

type Client struct {
	Requester rpcClient.EndpointRequester
}

func NewClient(uri string) *Client {
	return &Client{Requester: rpcClient.NewEndpointRequester(uri)}
}

// GetBlockByHeight returns the block bytes at the given height.
func (c *Client) GetBlockByHeight(ctx context.Context, height uint64, options ...rpcClient.Option) ([]byte, error) {
	res := &GetBlockByHeightResponse{}
	if err := c.Requester.SendRequest(ctx, "blocks.getBlockByHeight", &GetBlockByHeightArgs{Height: height}, res, options...); err != nil {
		return nil, err
	}

	blockBytes, err := hex.DecodeString(res.Block)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block %d retrieved via GetBlockByHeight: %w", height, err)
	}
	return blockBytes, nil
}

func (c *Client) GetMaxHeight(ctx context.Context, options ...rpcClient.Option) (uint64, error) {
	res := &GetMaxHeightResponse{}
	if err := c.Requester.SendRequest(ctx, "blocks.getMaxHeight", &struct{}{}, res, options...); err != nil {
		return 0, err
	}
	return res.Height, nil
}
