package cmd

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
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
	rootCmd.PersistentFlags().StringP("chain", "c", "mainnet", "The chain to update from")
	rootCmd.PersistentFlags().StringP("fmt", "x", "txt", "The format of the data files to process")
	rootCmd.MarkPersistentFlagRequired("folder")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	cobra.EnableCommandSorting = false
}

// getFolderAndChain returns the folder to process and the chain from the rootCmd. Note this function
// calls Fatal on error, therefore it does not return an error
func getOptions() (string, string, string) {
	folder, err := rootCmd.PersistentFlags().GetString("folder")
	if err != nil {
		log.Fatal(err)
	}
	cwd, _ := os.Getwd()
	folderPath := path.Join(cwd, folder)
	if len(folder) == 0 || !file.FolderExists(folderPath) {
		log.Fatal("You must provide a folder: ", folder)
	}

	chain, err := rootCmd.PersistentFlags().GetString("chain")
	if err != nil {
		log.Fatal(err)
	}
	err = validate.ValidateEnum("chain", chain, "[mainnet|gnosis]")
	if err != nil {
		log.Fatal(err)
	}

	format, err := rootCmd.PersistentFlags().GetString("fmt")
	if err != nil {
		log.Fatal(err)
	}
	if len(format) > 0 {
		err = validate.ValidateEnum("fmt", format, "[txt|csv|json]")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		format = "txt"
	}

	return folder, chain, format
}

var dataTypes = []string{
	"apps",
	"txs",
	"logs",
	"neighbors",
	// "neighbors/adjacencies",
	"statements",
	// "statements/balances",
	// "statements/tx_counts",
}

func colorHelp(helpIn string) string {
	ret := strings.Replace(helpIn, "{", colors.Cyan, -1)
	ret = strings.Replace(ret, "}", colors.Off, -1)
	return ret
}
