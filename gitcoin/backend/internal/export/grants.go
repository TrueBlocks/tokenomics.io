// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package exportPkg

// import (
// 	// "fmt"
// 	// "sort"

// 	// "github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/progress"
// 	// "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
// 	// "github.com/spf13/cobra"
// )

// type ProcessOptions struct {
// 	stats    bool
// 	grants   bool
// 	monitors bool
// }

// var opts ProcessOptions

// func RunE(cmd *cobra.Command, args []string) {

// 	if len(args) > 0 {
// 		if args[0] == "grants" {
// 			opts.grants = true
// 		} else {
// 			// default to monitors
// 			opts.monitors = true
// 		}

// 	}

// 	var summaries GrantSummaries

// 	progressChannel := progress.MakeChan()
// 	defer close(progressChannel)

// 	go ProcessGrants(progressChannel)

// 	var grantsDone uint
// 	for event := range progressChannel {
// 		grant, ok := event.Payload.(*Grant)
// 		var fileName string
// 		if ok {
// 			fileName = grant.Title
// 		}

// 		if event.Event == progress.Finished {
// 			logger.Log(logger.Info, grantsDone, "grant(s) were processed")
// 			break
// 		}

// 		switch event.Event {

// 		case progress.Error:
// 			logger.Log(logger.Error, event.Message)

// 		case progress.Start:
// 			logger.Log(logger.Info, event.Message, fileName)

// 		case progress.Update:
// 			logger.Log(logger.Info, event.Message, fileName)

// 		case progress.Done:
// 			if ok {
// 				if grant.Logo == nil {
// 					str := ""
// 					grant.Logo = &str
// 				}
// 				var summary GrantSummary
// 				summary.FromGrant(grant)
// 				summaries = append(summaries, summary)
// 				// fmt.Printf("%04d\t%-42s\t%s\t%s\t%s\n", grant.Id, strings.ToLower(grant.AdminAddress), grant.Slug, grant.Title, *grant.Logo)
// 			}
// 			grantsDone++

// 		default:
// 			logger.Log(logger.Info, event.Message, fileName)
// 		}
// 	}

// 	sort.Slice(summaries, func(i, j int) bool {
// 		iAddr := summaries[i]
// 		jAddr := summaries[j]
// 		return iAddr.Address < jAddr.Address
// 	})

// 	fmt.Printf("[")
// 	for i, summary := range summaries {
// 		if summary.Address == "0x0" {
// 			continue
// 		}
// 		if summary.Id == 1921 {
// 			continue
// 		}
// 		if i > 0 {
// 			fmt.Printf(",")
// 		}
// 		summary.Key = uint64(i)
// 		fmt.Printf("%s", summary.ToJSON())
// 	}
// 	fmt.Printf("]")
// }

// func ProcessGrants(progressChannel chan<- *progress.Progress) {
// 	progressChannel <- progress.StartMsg("Processing grants", nil)

// 	if opts.stats {
// 		fmt.Printf("GrantID\tAddress\tFileSize\tnAppearances\tFirst App\tLatest App\tBlockRange\tAgeInBlocks\n")
// 	}

// 	for i := 0; i < 4000; i++ {
// 		grantId := fmt.Sprintf("../data/raw/%04d.json", i)
// 		var grant Grant
// 		err := grant.GetGrant(grantId)
// 		if err != nil {
// 			progressChannel <- progress.ErrorMsg("Error processing grant "+grantId+" "+err.Error(), nil)
// 		} else {
// 			progressChannel <- progress.UpdateMsg("Processing grant id: "+grantId, nil)
// 			if opts.stats {
// 				monitor, err := GetMonitorStats(grantId, grant.AdminAddress)
// 				if err != nil {
// 					progressChannel <- progress.ErrorMsg("Error processing grant "+grantId+" "+err.Error(), nil)
// 				} else {
// 					fmt.Printf("%d\t%s\t%d\t%d\t%d.%d\t%d.%d\t%d\t%d\n", grant.Id, grant.AdminAddress, monitor.Size, monitor.Count, monitor.First.Bn, monitor.First.TxId, monitor.Last.Bn, monitor.Last.TxId, monitor.Range, monitor.Age)
// 				}
// 			}
// 			progressChannel <- progress.DoneMsg("----------> "+grantId, grant)
// 		}
// 	}

// 	progressChannel <- progress.FinishedMsg("Finished", nil)
// }
