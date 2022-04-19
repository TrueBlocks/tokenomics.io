/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/types"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		folder := "gitcoin"
		chain := "mainnet"

		grants := []types.Grant{}

		lines := file.AsciiFileToLines(folder + "/addresses.txt")
		for i, line := range lines {
			parts := strings.Split(line, "\t")
			if len(parts) < 2 {
				parts = append(parts, fmt.Sprintf("%04d", i))
			}
			if len(parts) < 3 {
				parts = append(parts, "Name "+parts[1])
			}
			if len(parts) < 4 {
				parts = append(parts, "Active")
			}
			if len(parts) < 5 {
				v := "false"
				if strings.Contains(parts[0], "Core") {
					v = "true"
				}
				parts = append(parts, v)
			}

			grant := types.Grant{
				GrantId:     parts[1],
				Address:     parts[0],
				Name:        parts[2],
				IsActive:    parts[3] == "Active",
				Core:        parts[4] == "true",
				LastUpdated: time.Now().Unix(),
			}

			chainData := types.Chain{Name: "mainnet"}
			var err error
			chainData.Counts, err = file.LineCounts(folder, chain, grant.Address)
			if err != nil {
				log.Fatal(err)
			}
			chainData.Types = chainData.Counts.Types()
			grant.Chains = append(grant.Chains, chainData)

			grants = append(grants, grant)
		}

		for _, grant := range grants {
			if grant.Chains[0].HasRecords() {
				fmt.Println(grant)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
