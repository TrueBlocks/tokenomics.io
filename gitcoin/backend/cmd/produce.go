// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/TrueBlocks/tokenomics/gitcoin/backend/grants"
	"github.com/TrueBlocks/tokenomics/gitcoin/backend/progress"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Produce,
}

func init() {
	rootCmd.AddCommand(produceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// produceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// produceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Produce(cmd *cobra.Command, args []string) {

	var summaries grants.GrantSummaries

	progressChannel := progress.MakeChan()
	defer close(progressChannel)

	go ProcessGrants(progressChannel)

	var grantsDone uint
	for event := range progressChannel {
		grant, ok := event.Payload.(*grants.Grant)
		var fileName string
		if ok {
			fileName = grant.Title
		}

		if event.Event == progress.Finished {
			logger.Log(logger.Info, grantsDone, "grant(s) were processed")
			break
		}

		switch event.Event {

		case progress.Error:
			logger.Log(logger.Error, event.Message)

		case progress.Start:
			logger.Log(logger.Info, event.Message, fileName)

		case progress.Update:
			logger.Log(logger.Info, event.Message, fileName)

		case progress.Done:
			if ok {
				if grant.Logo == nil {
					str := ""
					grant.Logo = &str
				}
				var summary grants.GrantSummary
				summary.FromGrant(grant)
				summaries = append(summaries, summary)
				// fmt.Printf("%04d\t%-42s\t%s\t%s\t%s\n", grant.Id, strings.ToLower(grant.AdminAddress), grant.Slug, grant.Title, *grant.Logo)
			}
			grantsDone++

		default:
			logger.Log(logger.Info, event.Message, fileName)
		}
	}

	sort.Slice(summaries, func(i, j int) bool {
		iAddr := summaries[i]
		jAddr := summaries[j]
		return iAddr.Address < jAddr.Address
	})

	fmt.Printf("[")
	for i, summary := range summaries {
		if summary.Address == "0x0" {
			continue
		}
		if summary.Id == 1921 {
			continue
		}
		if i > 0 {
			fmt.Printf(",")
		}
		summary.Key = uint64(i)
		fmt.Printf("%s", summary.ToJSON())
	}
	fmt.Printf("]")
}

func ProcessGrants(progressChannel chan<- *progress.Progress) {
	progressChannel <- progress.StartMsg("I am just getting started", nil)

	for i := 0; i < 4000; i++ {
		grantId := fmt.Sprintf("../data/raw/%04d.json", i)
		progressChannel <- progress.UpdateMsg("Processing grant id: "+grantId, nil)

		jsonFile, err := os.Open(grantId)
		if err != nil {
			progressChannel <- progress.ErrorMsg("Could not open file for grant id: "+grantId, nil)
		}
		byteData, _ := ioutil.ReadAll(jsonFile)
		jsonFile.Close()

		// data on disc is stored as an array, but only contains a single grant
		var tmpArray grants.Grants
		err = json.Unmarshal(byteData, &tmpArray)
		if err != nil {
			progressChannel <- progress.ErrorMsg("Could not Unmarshal grant id: "+grantId, nil)

		} else {
			// many of the files contain an empty array, only process if not empty
			if len(tmpArray) > 0 {
				progressChannel <- progress.DoneMsg("----------> "+grantId, &tmpArray[0])
			}
		}
	}

	progressChannel <- progress.FinishedMsg("Finished", nil)
}
