/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tokenomics.io",
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
	rootCmd.PersistentFlags().StringP("chain", "c", "mainnet", "The chain to update from (default 'mainnet'")
	rootCmd.PersistentFlags().StringP("folder", "f", "", "The local folder to process (required)")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tokenomics.io.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// getFolderAndChain returns the folder to process and the chain from the rootCmd. Note this function
// calls Fatal on error, therefore it does not return an error
func getFolderAndChain() (string, string) {
	folder, err := rootCmd.PersistentFlags().GetString("folder")
	if err != nil {
		log.Fatal(err)
	}
	if len(folder) == 0 || !file.FolderExists(folder) {
		log.Fatal("You must provide a folder")
	}
	chain, err := rootCmd.PersistentFlags().GetString("chain")
	if err != nil {
		log.Fatal(err)
	}
	if chain != "gnosis" && chain != "mainnet" {
		log.Fatal("only gnosis and mainnet are currently supported")
	}
	return folder, chain
}
