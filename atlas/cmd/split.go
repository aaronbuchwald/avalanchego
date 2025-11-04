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

	atlascontext "github.com/ava-labs/avalanchego/atlas/context"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	sourceStateDirFlag = "source-state-dir"
	targetStateDirFlag = "target-state-dir"
	heightFlag         = "height"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a shard at a given height to a new target state directory",
	RunE:  runSplit,
}

func init() {
	rootCmd.AddCommand(splitCmd)
	splitCmd.PersistentFlags().String(logLevelFlag, "info", "The log level to use")
	registerSplitFlags(splitCmd)
}

func registerSplitFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(sourceStateDirFlag, "", "The directory to store the source state of the shard")
	cmd.PersistentFlags().String(targetStateDirFlag, "", "The directory to store the target state of the shard")
	cmd.PersistentFlags().String(heightFlag, "", "The height to split the shard at (supports suffixes: k, m, b)")
}

func getSplitFlags(cmd *cobra.Command) (sourceStateDir string, targetStateDir string, height uint64, err error) {
	sourceStateDir, err = cmd.PersistentFlags().GetString(sourceStateDirFlag)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get source state directory: %w", err)
	}
	targetStateDir, err = cmd.PersistentFlags().GetString(targetStateDirFlag)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get target state directory: %w", err)
	}
	heightStr, err := cmd.PersistentFlags().GetString(heightFlag)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to get height: %w", err)
	}
	height, err = parseHeight(heightStr)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to parse height: %w", err)
	}
	return sourceStateDir, targetStateDir, height, nil
}

func runSplit(cmd *cobra.Command, args []string) error {
	sourceStateDir, targetStateDir, height, err := getSplitFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to get split flags: %w", err)
	}
	if err := initLogger(cmd); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	ctx, cancel := atlascontext.WithDefaultSignals(context.Background())
	defer cancel()

	sourceShard, err := shardFactory.New(ctx, log, sourceStateDir)
	if err != nil {
		return fmt.Errorf("failed to create source shard: %w", err)
	}
	defer sourceShard.Shutdown(ctx)

	log.Info("Performing split at height",
		zap.Uint64("height", height),
		zap.String("sourceStateDir", sourceStateDir),
		zap.String("targetStateDir", targetStateDir),
	)

	return sourceShard.SplitAtHeight(ctx, height, targetStateDir)
}
