// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ava-labs/avalanchego/atlas/evm"
	"github.com/ava-labs/avalanchego/atlas/shard"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	shardFactory shard.ShardFactory = evm.EVMShardFactory
	log          logging.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atlas",
	Short: "Atlas is a cloud-native archival blockchain service",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.atlas.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".atlas" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".atlas")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func initLogger(cmd *cobra.Command) error {
	logLevelStr, err := cmd.PersistentFlags().GetString(logLevelFlag)
	if err != nil {
		return fmt.Errorf("failed to get log level: %w", err)
	}
	logLevel, err := logging.ToLevel(logLevelStr)
	if err != nil {
		return fmt.Errorf("failed to parse log level: %w", err)
	}
	log = logging.NewLogger(
		"atlas",
		logging.NewWrappedCore(
			logLevel,
			os.Stdout,
			logging.Colors.ConsoleEncoder(),
		),
	)
	return nil
}

// contextWithDefaultSignals returns a context and cancellation function where the context is cancelled
// when SIGINT or SIGTERM are received.
func contextWithDefaultSignals(ctx context.Context) (context.Context, context.CancelFunc) {
	return contextWithSignals(ctx, syscall.SIGINT, syscall.SIGTERM)
}

// contextWithSignals returns a context and cancellation function where the context is cancelled
// when any of the provided signals are received.
func contextWithSignals(ctx context.Context, sig ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)

	go func() {
		defer cancel()
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			return
		}
	}()

	return ctx, cancel
}
