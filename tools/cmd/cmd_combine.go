package cmd

import (
	"github.com/TrueBlocks/tokenomics.io/tools/internal"
	"github.com/spf13/cobra"
)

const useCombine = "combine"
const shortCombine = "Combines per address data for each data type into a single file for that type"
const longCombine = `
This subcommand processes the addresses in the given {--folder} and {--chain}
and creates compressed tar files ({.tar.gz}) for each data type. The results
are placed in {./<folder>/exports/<chain>/combined/}.
`

var combineCmd = &cobra.Command{
	Use:   useCombine,
	Short: shortCombine,
	Long:  internal.ColorHelp(longCombine),
	RunE:  internal.RunCombine,
}

func init() {
	rootCmd.AddCommand(combineCmd)
}
