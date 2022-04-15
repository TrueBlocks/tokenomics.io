/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	tokenomics "github.com/TrueBlocks/tokenomics.io/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FlagFileConfig struct {
	InputDirectory string
	OutputFile     string
}

var flagToFileConfig = map[string]FlagFileConfig{
	"appearances": {
		InputDirectory: "./apps", // TODO: change it to configurable source directory
		OutputFile:     "./combined/apps.csv",
	},
}

// combineCmd represents the combine command
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("combine called")

		addressesFilePath := args[0]

		if cmd.Flags().NFlag() == 0 {
			return errors.New("at least one flag required")
		}

		flagsSet := []string{}
		cmd.Flags().Visit(func(f *pflag.Flag) {
			flagsSet = append(flagsSet, f.Name)
		})

		reader, err := tokenomics.ReadGrants(addressesFilePath)
		if err != nil {
			return err
		}

		flagToOutputFiles := map[string]*os.File{}
		for _, flagName := range flagsSet {
			fileConfig, ok := flagToFileConfig[flagName]
			if !ok {
				return fmt.Errorf("missing file configuration for flag: %s", flagName)
			}
			output, err := os.OpenFile(fileConfig.OutputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				return err
			}
			flagToOutputFiles[flagName] = output
			// runs at the end of the function
			defer output.Close()
		}

		for {
			grant, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			for _, flagName := range flagsSet {
				fileConfig := flagToFileConfig[flagName]

				input, err := os.ReadFile(path.Join(fileConfig.InputDirectory, grant.Address) + ".csv")
				if err != nil {
					return err
				}

				output := flagToOutputFiles[flagName]
				_, err = output.Write(input)
				if err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// combineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// combineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	combineCmd.Flags().BoolP("appearances", "a", false, "Output appearances file")
}
