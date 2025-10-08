// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package shard

import (
	"context"

	pb "github.com/ava-labs/avalanchego/atlas/proto/pb/writeshard"
)

type WriteShardServer struct {
	pb.UnimplementedWriteShardServer

	shard WriteShard
}

func NewGRPCShardServer(shard WriteShard) *WriteShardServer {
	return &WriteShardServer{
		shard: shard,
	}
}

func (s *WriteShardServer) ExecuteBlock(ctx context.Context, req *pb.ExecuteBlockRequest) (*pb.ExecuteBlockResponse, error) {
	if err := s.shard.ExecuteBlock(ctx, req.BlockBytes); err != nil {
		return nil, err
	}
	return &pb.ExecuteBlockResponse{}, nil
}
