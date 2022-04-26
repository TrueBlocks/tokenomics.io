package internal

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

// getFolderAndChain returns the folder to process and the chain from the rootCmd. Note this function
// calls Fatal on error, therefore it does not return an error
func getOptions(cmd *cobra.Command) (string, string, string) {
	folder, err := cmd.PersistentFlags().GetString("folder")
	if err != nil {
		log.Fatal(err)
	}
	cwd, _ := os.Getwd()
	folderPath := path.Join(cwd, folder)
	if len(folder) == 0 || !file.FolderExists(folderPath) {
		log.Fatal("You must provide a folder: ", folder)
	}

	chain, err := cmd.PersistentFlags().GetString("chain")
	if err != nil {
		log.Fatal(err)
	}
	err = validate.ValidateEnum("chain", chain, "[mainnet|gnosis]")
	if err != nil {
		log.Fatal(err)
	}

	format, err := cmd.PersistentFlags().GetString("fmt")
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

func ColorHelp(helpIn string) string {
	ret := strings.Replace(helpIn, "{", colors.Cyan, -1)
	ret = strings.Replace(ret, "}", colors.Off, -1)
	return ret
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
