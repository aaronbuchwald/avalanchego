// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package shard

import (
	"context"
	"net"

	pb "github.com/ava-labs/avalanchego/atlas/proto/pb/writeshard"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/grpcutils"
	"github.com/ava-labs/libevm/log"
	"google.golang.org/grpc"
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

func ServeGRPCShard(ctx context.Context, listener net.Listener, shard WriteShard) {
	shardServer := NewGRPCShardServer(shard)
	grpcServer := grpc.NewServer()
	pb.RegisterWriteShardServer(grpcServer, shardServer)

	go func() {
		defer func() {
			grpcServer.GracefulStop()
			log.Info("gRPC shard server completed graceful shutdown")
		}()

		<-ctx.Done()
		log.Info("Shutting down gRPC shard server...")
	}()

	grpcutils.Serve(listener, grpcServer)
}
