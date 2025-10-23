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
	"strconv"
	"strings"

	atlascontext "github.com/ava-labs/avalanchego/atlas/context"
	"github.com/ava-labs/avalanchego/atlas/evm"
	atlashttp "github.com/ava-labs/avalanchego/atlas/http"
	"github.com/spf13/cobra"
)

const (
	shardsFlag = "shards"
)

// routerCmd represents the router command
var routerCmd = &cobra.Command{
	Use:   "router",
	Short: "Runs a router that dispatches requests to the appropriate shard",
	RunE:  runRouter,
}

func init() {
	rootCmd.AddCommand(routerCmd)

	routerCmd.PersistentFlags().StringArray(shardsFlag, []string{}, "The addresses of the shards to route to")
	routerCmd.PersistentFlags().Int(portFlag, 8080, "The port to listen on")
	routerCmd.PersistentFlags().String(logLevelFlag, "info", "The log level to use")
}

// parseHeight converts a height string to uint64, supporting k/m/b suffixes
// Examples: "10000000", "10m", "1.5m", "20k", "5b"
func parseHeight(s string) (uint64, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return 0, fmt.Errorf("empty height value")
	}

	multiplier := uint64(1)
	if strings.HasSuffix(s, "k") {
		multiplier = 1_000
		s = strings.TrimSuffix(s, "k")
	} else if strings.HasSuffix(s, "m") {
		multiplier = 1_000_000
		s = strings.TrimSuffix(s, "m")
	} else if strings.HasSuffix(s, "b") {
		multiplier = 1_000_000_000
		s = strings.TrimSuffix(s, "b")
	}

	// Try parsing as float to support decimals like "1.5m"
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid height value %q: %w", s, err)
	}

	if val < 0 {
		return 0, fmt.Errorf("height cannot be negative: %q", s)
	}

	result := uint64(val * float64(multiplier))
	return result, nil
}

// parseHeightRange parses a height range string into start and end values
// Examples: "0-10m", "10m-20m", "1000-2000000"
func parseHeightRange(s string) (start, end uint64, err error) {
	s = strings.TrimSpace(s)

	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range format %q: expected start-end", s)
	}

	start, err = parseHeight(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start height in range %q: %w", s, err)
	}

	end, err = parseHeight(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid end height in range %q: %w", s, err)
	}

	if start > end {
		return 0, 0, fmt.Errorf("start height %d cannot be greater than end height %d in range %q", start, end, s)
	}

	return start, end, nil
}

func runRouter(cmd *cobra.Command, args []string) error {
	port, err := cmd.PersistentFlags().GetInt(portFlag)
	if err != nil {
		return fmt.Errorf("failed to get port: %w", err)
	}

	if err := initLogger(cmd); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	shardSpecs, _ := cmd.Flags().GetStringArray(shardsFlag)
	shards := make([]*evm.APIShard, 0, len(shardSpecs))
	for _, spec := range shardSpecs {
		parts := strings.SplitN(spec, ",", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid shard spec %q: expected format endpoint,start-end", spec)
		}
		start, end, err := parseHeightRange(parts[1])
		if err != nil {
			return fmt.Errorf("invalid height range in %q: %w", spec, err)
		}
		shards = append(shards, &evm.APIShard{
			Endpoint: parts[0],
			Start:    start,
			End:      end,
		})
	}

	// TODO: define constructor of base level command that takes in required factories, so the entire CLI can be constructed for
	// a specific set of concrete implementations.
	router := evm.NewRouter(shards)

	ctx, cancel := atlascontext.WithDefaultSignals(context.Background())
	defer cancel()

	return atlashttp.ServeWithContext(
		ctx,
		router,
		atlashttp.WithLogger(log),
		atlashttp.WithPort(port),
	)
}
