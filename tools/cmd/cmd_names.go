package cmd

import (
	"github.com/TrueBlocks/tokenomics.io/tools/internal"
	"github.com/spf13/cobra"
)

const useNames = "names"
const shortNames = "generate an addName script to export addresses.tsv file to chifra names database"
const longNames = `
This routine takes the addresses.tsv file and generates an addNames script and/or takes
the results of "chifra names" and generates an addresses.tsv file.
`

var namesCmd = &cobra.Command{
	Use:   useNames,
	Short: shortNames,
	Long:  internal.ColorHelp(longNames),
	RunE:  internal.RunNames,
}

func init() {
	rootCmd.AddCommand(namesCmd)
	namesCmd.Flags().StringP("tag", "t", "", "Add a tag to each address")
}
