/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/monitor"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/index"
	tslibPkg "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/tslib"
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
			slug := "https://gitcoin.co/grants/" + parts[1] + "/" + strings.Replace(strings.ToLower(parts[2]), ".", "", -1)
			slug = strings.Replace(slug, " ", "-", -1)
			parts = append(parts, slug)

			grant := types.Grant{
				GrantId:  parts[1],
				Address:  strings.ToLower(parts[0]),
				Name:     parts[2],
				IsActive: parts[3] == "Active" || parts[3] == "true",
				Core:     parts[4] == "true",
				// TODO: BOGUS - fix this in production
				LastUpdated: 0, // time.Now().Unix(),
				Slug:        parts[5],
			}

			chainData := types.Chain{ChainName: "mainnet"}
			var err error
			chainData.Counts, err = LineCounts(folder, chain, grant.Address)
			if err != nil {
				log.Fatal(err)
			}

			mon := monitor.NewMonitor(chain, grant.Address, false)
			chainData.FileSize = file.FileSize(mon.Path())
			if file.FileExists(mon.Path()) {
				err = mon.ReadHeader()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				apps := make([]index.AppearanceRecord, mon.Count(), mon.Count())
				err = mon.ReadAppearances(&apps)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if len(apps) > 0 {
					chainData.FirstApp.Bn = int(apps[0].BlockNumber)
					chainData.FirstApp.TxId = int(apps[0].TransactionId)
					t, _ := tslibPkg.TsFromBn(chain, uint64(apps[0].BlockNumber))
					chainData.FirstApp.Timestamp = int(t)
					chainData.FirstApp.Date, _ = tslibPkg.DateFromTs(uint64(chainData.FirstApp.Timestamp))
					chainData.LatestApp.Bn = int(apps[len(apps)-1].BlockNumber)
					chainData.LatestApp.TxId = int(apps[len(apps)-1].TransactionId)
					t, _ = tslibPkg.TsFromBn(chain, uint64(apps[len(apps)-1].BlockNumber))
					chainData.LatestApp.Timestamp = int(t)
					chainData.LatestApp.Date, _ = tslibPkg.DateFromTs(uint64(chainData.LatestApp.Timestamp))
					chainData.BlockRange = chainData.LatestApp.Bn - chainData.FirstApp.Bn + 1
				}
			}
			mon.Close()
			chainData.Types = chainData.Counts.Types()
			grant.Chains = append(grant.Chains, chainData)

			grants = append(grants, grant)
		}

		fmt.Println("[")
		for i, grant := range grants {
			if grant.Chains[0].HasRecords() {
				if i > 0 {
					fmt.Println(",")
				}
				str := grant.String()
				fmt.Println(str)
			}
		}
		fmt.Println("]")
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

func LineCounts(folder, chain, addr string) (types.Counts, error) {
	if !strings.HasSuffix(folder, "/") {
		folder += "/"
	}
	base := "./" + folder + "exports/" + chain
	if !file.FolderExists(base) {
		fmt.Println("SOMSOMTEOMTOME")
		return types.Counts{}, fmt.Errorf("data folder (%s) not found", base)
	}
	counts := types.Counts{}
	counts.Appearances, _ = file.LineCount(folder+"exports/"+chain+"/apps/"+addr+".csv", true)
	counts.Neighbors, _ = file.LineCount(folder+"exports/"+chain+"/neighbors/"+addr+".csv", true)
	counts.Logs, _ = file.LineCount(folder+"exports/"+chain+"/logs/"+addr+".csv", true)
	counts.Txs, _ = file.LineCount(folder+"exports/"+chain+"/txs/"+addr+".csv", true)
	counts.Statements, _ = file.LineCount(folder+"exports/"+chain+"/statements/"+addr+".csv", true)
	return counts, nil
}
