package cmd

import (
	"io"
	"log"
	"os"
	"path"
	"sync"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/spf13/cobra"
)

// combineCmd represents the combine command
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Combines per address data into single file",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, chain := getFolderAndChain()

		// Define where to find addresses file
		grantReader, err := tokenomics.ReadGrants("./addresses.txt")
		if err != nil {
			log.Fatal(err)
		}

		// Create a map of open output files so we don't have to re-open
		// and close for each grant. Use archiveInputDirs
		flagToOutputFiles := map[string]*os.File{}
		for _, dir := range archiveInputDirs {
			outputPath := path.Join("./exports", chain, "combined", dir) + ".csv"
			output, err := os.Create(outputPath)
			if err != nil {
				log.Fatal(err)
			}
			flagToOutputFiles[dir] = output
			defer output.Close() // when the function exits
		}

		for {
			// Read one address
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

			log.Println("Combining", grant.Address)

			// For each flag, we will combine the single address file. We want to
			// do it concurrently, because opening input file can take some time.
			var wg sync.WaitGroup
			for _, dir := range archiveInputDirs {
				wg.Add(1)
				// This go routine contains main logic
				go func() {
					defer wg.Done()

					inputPath := path.Join("./exports", chain, dir, grant.Address) + ".csv"
					input, err := os.ReadFile(inputPath)
					if err != nil {
						log.Println(err)
					} else {
						outputFile := flagToOutputFiles[dir]
						_, err = outputFile.Write(input)
						if err != nil {
							log.Fatal(err)
						}
					}
				}()
				wg.Wait()
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)
}
