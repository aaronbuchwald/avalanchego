// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	flaghelpers "github.com/ava-labs/avalanchego/cmd/flags"
)

const (
	LogLevelKey        = "log-level"
	SnowTypeKey        = "snow"
	KKey               = "k"
	AlphaConfidenceKey = "alpha-confidence-key"
	AlphaPreferenceKey = "alpha-preference"
	BetaVirtuousKey    = "beta-virtuous"
	BetaRogueKey       = "beta-rogue"
	NKey               = "n"
	BlueKey            = "blue"
	ByzKey             = "byz"
	OutputFileKey      = "output-file"
)

func BuildViper(args []string) (*viper.Viper, error) {
	return flaghelpers.BuildViper("network", func(fs *pflag.FlagSet) {
		fs.String(LogLevelKey, "info", "Specify the log level")
		fs.String(SnowTypeKey, "snowflake", "Specify the consensus implementation to use (Snowball/Snowflake).")
		fs.Int(KKey, 20, "Specify the poll sample size K.")
		fs.Int(AlphaConfidenceKey, 15, "Specify the alpha confidence threshold required to record a successful poll.")
		fs.Int(AlphaPreferenceKey, 11, "Specify the alpha preference threshold required to record a successful preference poll.")
		fs.Int(BetaVirtuousKey, 20, "Specify the number of consecutive successful polls required to commit a virtuous color.")
		fs.Int(BetaRogueKey, 20, "Specify the number of consecutive successful polls required to commit a rogue color.")
		fs.Int(NKey, 500, "Specify the total number of nodes to include in the simulation.")
		fs.Float64(BlueKey, 0.5, "Specify the portion of the virtuous nodes to initialize to blue.")
		fs.Float64(ByzKey, 0.0, "Specify the byzantine portion of the network.")
		fs.String(OutputFileKey, "", "Specify the location to write results to a CSV file.")
	}, args)
}
