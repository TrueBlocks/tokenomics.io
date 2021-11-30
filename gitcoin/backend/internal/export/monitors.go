// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package exportPkg

import (
	"fmt"

	"github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/progress"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/spf13/cobra"
)

type ProcessOptions struct {
	Stats  bool
	Format string
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

func ProcessGrants(progressChannel chan<- *progress.Progress) {
	progressChannel <- progress.StartMsg("Processing grants", nil)

	if Options.Stats {
		if Options.Format != "json" {
			if Options.Format == "txt" {
				fmt.Printf("GrantID\tAddress\tFileSize\tnAppearances\tFirst App\tLatest App\tBlockRange\tAgeInBlocks\n")

			} else {
				fmt.Printf("GrantID,Address,FileSize,nAppearances,First App,Latest App,BlockRange,AgeInBlocks\n")
			}
		} else {
			fmt.Printf("[\n")
		}
	}

	max := 20 // 4000
	first := true
	for i := 0; i < max; i++ {
		grantId := fmt.Sprintf("../data/raw/%04d.json", i)
		var grant Grant
		err := grant.GetGrant(grantId)
		if err != nil {
			progressChannel <- progress.ErrorMsg("Error processing grant "+grantId+" "+err.Error(), nil)
		} else {
			progressChannel <- progress.UpdateMsg("Processing grant id: "+grantId, nil)
			if Options.Stats {
				monitor, err := GetMonitorStats(grantId, grant.AdminAddress)
				if err != nil {
					progressChannel <- progress.ErrorMsg("Error processing grant "+grantId+" "+err.Error(), nil)
				} else {
					if Options.Format != "json" {
						if Options.Format == "txt" {
							fmt.Printf("%d\t%s\t%d\t%d\t%d.%d\t%d.%d\t%d\t%d\n", grant.Id, grant.AdminAddress, monitor.Size, monitor.Count, monitor.First.Bn, monitor.First.TxId, monitor.Latest.Bn, monitor.Latest.TxId, monitor.Range, monitor.Age)
						} else {
							fmt.Printf("%d,%s,%d,%d,%d.%d,%d.%d,%d,%d\n", grant.Id, grant.AdminAddress, monitor.Size, monitor.Count, monitor.First.Bn, monitor.First.TxId, monitor.Latest.Bn, monitor.Latest.TxId, monitor.Range, monitor.Age)
						}
					} else {
						if !first {
							fmt.Printf(",")
						}
						fmt.Printf("%s", monitor.ToJSON())
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
