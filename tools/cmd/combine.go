package cmd

import (
	"io"
	"log"
	"os"
	"path"
	"sync"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
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
	Long:  colorHelp(longCombine),
	RunE:  runCombine,
}

func init() {
	rootCmd.AddCommand(combineCmd)
}

func runCombine(cmd *cobra.Command, args []string) error {
	folder, chain, format := getOptions()

	addressFn := path.Join(folder, "./addresses.txt")
	if !file.FileExists(addressFn) {
		return validate.Usage("Cannot find address file {0}", addressFn)
	}

	// Define where to find addresses file
	grantReader, err := tokenomics.ReadGrants(addressFn)
	if err != nil {
		log.Fatal(err)
	}

	// Create a map of open output files so we don't have to re-open and close for each grant.
	typeToFileMap := map[string]*os.File{}
	for _, dataType := range dataTypes {
		outputPath := path.Join(folder, "./exports", chain, "combined", dataType) + "." + format
		output, err := os.Create(outputPath)
		if err != nil {
			log.Fatal(err)
		}
		typeToFileMap[dataType] = output
		defer output.Close() // note that this closes when the function goes out of scope, not this code block
	}

	headerWritten := make(map[string]bool, len(dataTypes))
	for {
		grant, err := grantReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !grant.IsValid {
			continue
		}

		log.Println("Combining data from", grant.Address)

		// For each data type, we combine the individual files for the current grant. We
		// do it concurrently, because opening input files takes time.
		var wg sync.WaitGroup
		for _, dataType := range dataTypes {
			wg.Add(1)
			go func() {
				defer wg.Done()

				inputPath := path.Join(folder, "./exports", chain, dataType, grant.Address) + "." + format
				outputFile := typeToFileMap[dataType]

				if file.FileExists(inputPath) {
					lines := file.AsciiFileToLines(inputPath)
					for i, line := range lines {
						line += "\n"
						if i == 0 {
							if !headerWritten[dataType] {
								if dataType != "apps" {
									if format == "txt" {
										line = "address\t" + line
									} else {
										line = "\"address\"," + line
									}
								}
								_, err = outputFile.WriteString(line)
								if err != nil {
									log.Fatal(err)
								}
								headerWritten[dataType] = true
							}
						} else {
							if dataType != "apps" {
								// Generally speaking, the per-address data files do not contain
								// the address itself (it's in the name of the file). We add that
								// here when we combine the data.
								// TODO: Should be just store the data in the file directly and remove this?
								if format == "txt" {
									line = grant.Address + "\t" + line
								} else {
									line = "\"" + grant.Address + "\"," + line
								}
							}
							_, err = outputFile.WriteString(line)
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				}
			}()
			wg.Wait()
		}
	}

	return nil
}
