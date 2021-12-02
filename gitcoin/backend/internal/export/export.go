// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package exportPkg

import (
	"fmt"
	"os"
	"strings"

	"github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/progress"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/spf13/cobra"
)

const pathToData = "../data/" // /Users/jrush/Development/tokenomics.io/gitcoin/data/"

type ProcessOptions struct {
	Stats   bool
	Scripts bool
	Format  string
}

var Options ProcessOptions

func RunE(cmd *cobra.Command, args []string) {

	progressChannel := progress.MakeChan()
	defer close(progressChannel)

	go ProcessGrants(progressChannel)

	var grantsDone uint
	for event := range progressChannel {
		grant, ok := event.Payload.(*Grant)
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
			grantsDone++

		default:
			logger.Log(logger.Info, event.Message, fileName)
		}
	}
}

type LastUpdate struct {
	Address string `json:"address"`
	Block   uint64 `json:"block"`
}

func ProcessGrants(progressChannel chan<- *progress.Progress) {
	var fileNames []string
	max := 4000
	for i := 0; i < max; i++ {
		fileNames = append(fileNames, fmt.Sprintf(pathToData+"raw/%04d.json", i))
	}
	for i := 0; i < 5; i++ {
		fileNames = append(fileNames, fmt.Sprintf(pathToData+"raw/core-%04d.json", i))
	}

	msg := fmt.Sprintf("Processing %d grants", len(fileNames))
	progressChannel <- progress.StartMsg(msg, nil)

	if Options.Format == "json" {
		fmt.Printf("[\n")

	} else {
		header := "GrantID\tAddress\tFileSize\tnAppearances\tFirst App\tLatest App\tBlockRange\tAgeInBlocks\n"
		if Options.Scripts {
			header = "Address\tLastUpdate\n"
		}
		if Options.Format == "csv" {
			header = strings.Replace(header, "\t", ",", -1)
		}
		fmt.Printf("%s", header)
	}

	first := true
	for _, grantId := range fileNames {
		fileStat, err := os.Stat(grantId)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if fileStat.Size() == 3 {
			continue
		}

		var grant Grant
		err = grant.GetGrant(grantId)
		if err != nil {
			progressChannel <- progress.ErrorMsg("Error processing grant "+grantId+" "+err.Error(), nil)
		} else {
			progressChannel <- progress.UpdateMsg("Processing grant id: "+grantId, nil)
			if Options.Scripts {
				var monitor Monitor
				monitor.Address = grant.AdminAddress
				var update LastUpdate
				update.Address = monitor.Address
				update.Block, _ = monitor.getLastUpdate()
				if len(update.Address) > 0 && update.Address != "0x0" {
					fmt.Printf("%s,%d\n", update.Address, update.Block)
				}
			} else if Options.Stats {
				monitor, err := GetMonitorStats(grantId, &grant)
				if err != nil {
					progressChannel <- progress.ErrorMsg("Error processing grant "+grantId+" "+err.Error(), nil)
				} else {
					grant.Monitor = *monitor
					if Options.Format != "json" {
						if Options.Format == "txt" {
							fmt.Printf("%d\t%s\t%d\t%d\t%d.%d\t%d.%d\t%d\n", grant.Id, grant.AdminAddress, grant.Monitor.Size, grant.Monitor.Count, grant.Monitor.First.Bn, grant.Monitor.First.TxId, grant.Monitor.Latest.Bn, grant.Monitor.Latest.TxId, grant.Monitor.Range)
						} else {
							fmt.Printf("%d,%s,%d,%d,%d.%d,%d.%d,%d\n", grant.Id, grant.AdminAddress, grant.Monitor.Size, grant.Monitor.Count, grant.Monitor.First.Bn, grant.Monitor.First.TxId, grant.Monitor.Latest.Bn, grant.Monitor.Latest.TxId, grant.Monitor.Range)
						}
					} else {
						if !first {
							fmt.Printf(",")
						}
						fmt.Printf("%s", grant.Monitor.ToJSON())
						first = false
					}
				}
			}
			progressChannel <- progress.DoneMsg("----------> "+grantId, grant)
		}
	}

	if Options.Format == "json" {
		fmt.Printf("]\n")
	}

	progressChannel <- progress.FinishedMsg("Finished", nil)
}
