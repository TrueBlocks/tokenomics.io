/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/spf13/cobra"
)

var archiveInputDirs = []string{
	"apps",
	"txs",
	"logs",
	"neighbors",
	// "neighbors/adjacencies",
	"statements",
	// "statements/balances",
	// "statements/tx_counts",
}

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Builds tar.gz files of each address in the pouch as well as per-type zips of all addresses",
	Long: `Assuming a file called ./addresses.txt in the local folder, this
tool reads that file and produces tar.gz files for each address containing all exported
data types. Additionally, it creates a single .tar.gz file for each data type containing
all addresses. Finally, it creates a single .tar.gz file containing all per-type .tar.gz
files so the entire database can be downloaded from a single file.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, chain := getFolderAndChain()

		// Define where to find addresses file
		grantReader, err := tokenomics.ReadGrants("./addresses.txt")
		if err != nil {
			log.Fatal(err)
		}

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

			log.Println("Compressing", grant.Address)

			tarFile, err := os.OpenFile(path.Join("./exports", chain, "zips", grant.Address+".tar"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatal(err)
			}
			archive := tar.NewWriter(tarFile)
			defer func() {
				if err := archive.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			for _, inputDir := range archiveInputDirs {
				log.Println("\t", inputDir)
				input, err := os.ReadFile(path.Join("./exports", chain, inputDir, grant.Address+".csv"))
				if err != nil {
					log.Println(err)
				} else {
					archive.Write(input)
				}
			}

			log.Println("Done", grant.Address)
			tarFile.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
