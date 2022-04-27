package cmd

import (
	"github.com/TrueBlocks/tokenomics.io/tools/internal"
	"github.com/spf13/cobra"
)

const useUpdate = "update"
const shortUpdate = "update each monitored group's theData.json file"
const longUpdate = `
This routine builds the 'database' for the front end ui for each monitored
group of addresses per chain. This means it reads the addresses.tsv file and processes
each address basically by counting how many of each of type of data is present.

The command can be run periodically (no more often that the scraper runs) by a cron
job for ecxample.
`

var updateCmd = &cobra.Command{
	Use:   useUpdate,
	Short: shortUpdate,
	Long:  internal.ColorHelp(longUpdate),
	RunE:  internal.RunUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
