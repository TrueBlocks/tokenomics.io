// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
package cmd

import (
	exportPkg "github.com/TrueBlocks/tokenomics.io/gitcoin/backend/internal/export"
	"github.com/spf13/cobra"
)

// exportCmd represents the monitors command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export information",
	Long:  `Export information`,
	Run:   exportPkg.RunE,
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().BoolVarP(&exportPkg.Options.Stats, "stats", "s", false, "produce only stats for the grants")
	exportCmd.Flags().StringVarP(&exportPkg.Options.Format, "fmt", "x", "csv", "format for the export")
	exportCmd.SetUsageTemplate(UsageText())
}
