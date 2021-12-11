// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
package cmd

import (
	names "github.com/TrueBlocks/tokenomics.io/gitcoin/backend/internal/names"
	"github.com/spf13/cobra"
)

// exportCmd represents the monitors command
var namesCmd = &cobra.Command{
	Use:   "names",
	Short: "Export names information",
	Long:  `Export names information`,
	Run:   names.RunE,
}

func init() {
	rootCmd.AddCommand(namesCmd)
	namesCmd.Flags().StringVarP(&names.Options.Format, "fmt", "x", "csv", "format for the export")
	namesCmd.SetUsageTemplate(UsageText())
}
