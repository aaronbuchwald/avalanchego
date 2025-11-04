/*
Copyright © 2025 Aaron Buchwald <aaron.buchwald56@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/atlas/avalanche/vms/evm"
	"github.com/ava-labs/avalanchego/atlas/blockdb"
	atlascontext "github.com/ava-labs/avalanchego/atlas/context"
	atlashttp "github.com/ava-labs/avalanchego/atlas/http"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	dbDirFlag           = "db-dir"
	wsEndpointFlag      = "ws-endpoint"
	initialMaxBlockFlag = "initial-max-block"
)

// blocksCmd represents the blocks command
var blocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "Serve a block database",
	RunE:  runBlocks,
}

func init() {
	rootCmd.AddCommand(blocksCmd)
}

func registerBlocksFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(dbDirFlag, "", "The directory to store the block database")
	cmd.PersistentFlags().Int(portFlag, 0, "The port to serve the block database on")
	cmd.PersistentFlags().String(wsEndpointFlag, "", "The websocket endpoint to listen for blocks from")
	cmd.PersistentFlags().Uint64(initialMaxBlockFlag, 0, "The initial maximum block to ingest")
}

func getBlocksFlags(cmd *cobra.Command) (dbDir string, port int, wsEndpoint string, initialMaxBlock uint64, err error) {
	dbDir, err = cmd.PersistentFlags().GetString(dbDirFlag)
	if err != nil {
		return "", 0, "", 0, fmt.Errorf("failed to get db directory: %w", err)
	}
	port, err = cmd.PersistentFlags().GetInt(portFlag)
	if err != nil {
		return "", 0, "", 0, fmt.Errorf("failed to get port: %w", err)
	}
	wsEndpoint, err = cmd.PersistentFlags().GetString(wsEndpointFlag)
	if err != nil {
		return "", 0, "", 0, fmt.Errorf("failed to get websocket endpoint: %w", err)
	}
	initialMaxBlock, err = cmd.PersistentFlags().GetUint64(initialMaxBlockFlag)
	if err != nil {
		return "", 0, "", 0, fmt.Errorf("failed to get initial maximum block: %w", err)
	}
	return dbDir, port, wsEndpoint, initialMaxBlock, nil
}

func runBlocks(cmd *cobra.Command, args []string) error {
	dbDir, port, wsEndpoint, initialMaxBlock, err := getBlocksFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to get blocks flags: %w", err)
	}

	opts := []blockdb.Option{}
	if initialMaxBlock > 0 {
		opts = append(opts, blockdb.WithMaxHeight(initialMaxBlock))
	}
	blockDB, err := blockdb.NewBlockDB(dbDir, opts...)
	if err != nil {
		return fmt.Errorf("failed to create block database: %w", err)
	}

	ctx, cancel := atlascontext.WithDefaultSignals(context.Background())
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		handler := blockdb.NewHTTPHandler(blockDB)

		return atlashttp.ServeWithContext(
			ctx,
			handler,
			atlashttp.WithLogger(log),
			atlashttp.WithPort(port),
		)
	})

	eg.Go(func() error {
		client, err := ethclient.DialContext(ctx, wsEndpoint)
		if err != nil {
			return fmt.Errorf("failed to dial websocket endpoint: %w", err)
		}
		defer client.Close()

		blockClient := evm.NewEVMBlockClient(client)
		blockResults, err := shard.CreateBlockStreamFromClient(ctx, blockClient, blockDB.GetMaxHeight())
		if err != nil {
			return fmt.Errorf("failed to create block result stream: %w", err)
		}
		return shard.IngestBlockStream(ctx, blockdb.NewBlockDBResultHandler(blockDB), blockResults)
	})

	return eg.Wait()
}
