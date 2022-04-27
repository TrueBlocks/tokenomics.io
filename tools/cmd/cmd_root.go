package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nomics",
	Short: "A tool to build datasets for tokenomics.io",
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
	rootCmd.PersistentFlags().StringP("folder", "f", "", "The local folder to process")
	rootCmd.PersistentFlags().StringArrayP("chain", "c", []string{"mainnet"}, "The chain to update from")
	rootCmd.PersistentFlags().StringP("fmt", "x", "txt", "The format of the data files to process")
	rootCmd.MarkPersistentFlagRequired("folder")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	cobra.EnableCommandSorting = false
}
