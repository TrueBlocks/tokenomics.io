/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"io"
	"log"
	"os"
	"path"
	"sync"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FlagFileConfig struct {
	InputDirectory string
	OutputFile     string
}

// TODO: CHAINS

var flagToFileConfig = map[string]FlagFileConfig{
	"appearances": {
		InputDirectory: "./apps",
		OutputFile:     "./combined/apps.csv",
	},
	"logs": {
		InputDirectory: "./logs",
		OutputFile:     "./combined/logs.csv",
	},
	"neighbors": {
		InputDirectory: "./neighbors",
		OutputFile:     "./combined/neighbors.csv",
	},
	"statements": {
		InputDirectory: "./statements",
		OutputFile:     "./combined/statements.csv",
	},
	"txs": {
		InputDirectory: "./txs",
		OutputFile:     "./combined/txs.csv",
	},
}

// combineCmd represents the combine command
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Combines per address data into single file",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Define where to find addresses file
		addressesFilePath := "./addresses.txt"

		if cmd.Flags().NFlag() == 0 {
			return errors.New("at least one flag required")
		}

		// We need to build a list of used flags so that we can match flags with
		// file configuration
		flagsSet := []string{}
		cmd.Flags().Visit(func(f *pflag.Flag) {
			flagsSet = append(flagsSet, f.Name)
		})

		// Create address file reader
		reader, err := tokenomics.ReadGrants(addressesFilePath)
		if err != nil {
			log.Fatal(err)
		}

		// We will create output files and keep them in a map, so that
		// we don't have to repeat open and close operations for each
		// address.
		flagToOutputFiles := map[string]*os.File{}
		for _, flagName := range flagsSet {
			fileConfig, ok := flagToFileConfig[flagName]
			if !ok {
				log.Fatalf("missing file configuration for flag: %s", flagName)
			}
			output, err := os.Create(fileConfig.OutputFile)
			if err != nil {
				log.Fatal(err)
			}
			flagToOutputFiles[flagName] = output
			// runs at the end of the function
			defer output.Close()
		}

		for {
			// Read one address
			grant, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			// For each flag, we will combine the single address file. We want to
			// do it concurrently, because opening input file can take some time.
			var wg sync.WaitGroup
			for _, flagName := range flagsSet {
				wg.Add(1)
				// This go routine contains main logic
				go func() {
					defer wg.Done()
					fileConfig := flagToFileConfig[flagName]

					input, err := os.ReadFile(path.Join(fileConfig.InputDirectory, grant.Address) + ".csv")
					if err != nil {
						log.Fatal(err)
					}

					output := flagToOutputFiles[flagName]
					_, err = output.Write(input)
					if err != nil {
						log.Fatal(err)
					}
				}()
				// Wait for all go routines for this address to finish
				wg.Wait()
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)

	combineCmd.Flags().BoolP("appearances", "a", false, "Output appearances file")
	combineCmd.Flags().BoolP("logs", "l", false, "Output logs file")
	combineCmd.Flags().BoolP("neighbors", "n", false, "Output neighbors file")
	combineCmd.Flags().BoolP("txs", "t", false, "Output txs file")
}
