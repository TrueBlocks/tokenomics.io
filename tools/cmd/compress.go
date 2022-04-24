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
	"logs/articulated",
	"neighbors",
	"neighbors/adjacencies",
	"statements",
	"statements/balances",
	"statements/tx_counts",
}

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define where to find addresses file
		addressesFilePath := "./addresses.txt"
		reader, err := tokenomics.ReadGrants(addressesFilePath)
		if err != nil {
			log.Fatal(err)
		}

		for {
			// Read one address
			grant, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Compressing", grant.Address)

			tarFile, err := os.OpenFile(path.Join("./zips", grant.Address+".tar"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatal(err)
			}
			defer tarFile.Close()
			archive := tar.NewWriter(tarFile)
			defer func() {
				if err := archive.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			for _, inputDir := range archiveInputDirs {
				log.Println("\t", inputDir)
				input, err := os.ReadFile(path.Join(inputDir, grant.Address+".csv"))
				if err != nil {
					log.Fatal(err)
				}
				archive.Write(input)
			}

			log.Println("Done", grant.Address)
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
