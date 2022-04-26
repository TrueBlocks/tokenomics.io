package cmd

import (
	"github.com/spf13/cobra"
)

const useCompress = "compress"
const shortCompress = "Builds tar.gz files of each address in the pouch as well as per-type .tar.gz files"
const longCompress = `
Assuming a file called ./addresses.txt in the local folder, this
tool reads that file and produces tar.gz files for each address containing all exported
data types. Additionally, it creates a single .tar.gz file for each data type containing
all addresses. Finally, it creates a single .tar.gz file containing all per-type .tar.gz
files so the entire database can be downloaded from a single file.
`

var compressCmd = &cobra.Command{
	Use:   useCompress,
	Short: shortCompress,
	Long:  colorHelp(longCompress),
	RunE:  runCompress,
}

func init() {
	rootCmd.AddCommand(compressCmd)
}
