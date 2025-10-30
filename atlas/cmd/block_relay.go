// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"

	atlascontext "github.com/ava-labs/avalanchego/atlas/context"
	pb "github.com/ava-labs/avalanchego/atlas/proto/pb/writeshard"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	websocketEndpointFlag  = "trusted-websocket-address"
	activeShardAddressFlag = "active-shard-address" // TODO: support multiple active shards
	lastAcceptedBlockFlag  = "start-block"
)

var blockRelayCmd = &cobra.Command{
	Use:   "block-relay",
	Short: "Relays blocks from a websocket endpoint to a CSV list of gRPC active shard services",
	RunE:  runBlockRelay,
}

func init() {
	rootCmd.AddCommand(blockRelayCmd)

	registerBlockRelayFlags(blockRelayCmd)
}

func registerBlockRelayFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(websocketEndpointFlag, "", "The websocket endpoint to listen for blocks from")
	cmd.PersistentFlags().String(activeShardAddressFlag, "", "The gRPC active shard address")
	cmd.PersistentFlags().Uint64(lastAcceptedBlockFlag, 0, "The last accepted block of the active shard to execute from (exclusive)")
}

func getBlockRelayFlags(cmd *cobra.Command) (websocketEndpoint string, activeShardAddress string, lastAcceptedBlock uint64, err error) {
	websocketEndpoint, err = cmd.PersistentFlags().GetString(websocketEndpointFlag)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get websocket endpoint: %w", err)
	}
	activeShardAddress, err = cmd.PersistentFlags().GetString(activeShardAddressFlag)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get active shard addresses: %w", err)
	}
	lastAcceptedBlock, err = cmd.PersistentFlags().GetUint64(lastAcceptedBlockFlag)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get last accepted block: %w", err)
	}

	return websocketEndpoint, activeShardAddress, lastAcceptedBlock, nil
}

func runBlockRelay(cmd *cobra.Command, args []string) error {
	websocketEndpoint, activeShardAddress, lastAcceptedBlock, err := getBlockRelayFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to get block relay flags: %w", err)
	}

	ctx, cancel := atlascontext.WithDefaultSignals(context.Background())
	defer cancel()

	return serveBlockRelay(ctx, lastAcceptedBlock, websocketEndpoint, activeShardAddress)
}

type BlockRelayHandler struct {
	shardClient pb.WriteShardClient
}

func NewBlockRelayHandler(shardClient pb.WriteShardClient) *BlockRelayHandler {
	return &BlockRelayHandler{
		shardClient: shardClient,
	}
}

func (h *BlockRelayHandler) HandleBlockResult(ctx context.Context, blockResult shard.BlockResult) error {
	_, err := h.shardClient.ExecuteBlock(ctx, &pb.ExecuteBlockRequest{BlockBytes: blockResult.Block})
	return err
}

func serveBlockRelay(ctx context.Context, lastAcceptedBlock uint64, websocketEndpoint string, activeShardAddress string) error {
	client, err := ethclient.DialContext(ctx, websocketEndpoint)
	if err != nil {
		return fmt.Errorf("failed to dial websocket endpoint: %w", err)
	}
	defer client.Close()

	grpcConn, err := grpc.NewClient(activeShardAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to create gRPC client: %w", err)
	}
	defer grpcConn.Close()

	shardClient := pb.NewWriteShardClient(grpcConn)

	blockResults, err := createBlockResultStreamFromClient(ctx, client, log, lastAcceptedBlock)
	if err != nil {
		return fmt.Errorf("failed to create block result stream: %w", err)
	}

	return shard.IngestBlockStream(ctx, NewBlockRelayHandler(shardClient), blockResults)
}
