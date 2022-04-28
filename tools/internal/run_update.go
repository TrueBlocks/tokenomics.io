package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"path"
	"sort"
	"strings"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/monitor"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/index"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpcClient"
	tslibPkg "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/tslib"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func RunUpdate(cmd *cobra.Command, args []string) error {
	folder, chains, format := getOptions(cmd.Parent())

	// ...list of existing grants
	grantMap, err := readExistingGrants(folder)
	if err != nil {
		log.Fatal(err)
	}

	addressFn := path.Join(folder, "./addresses.tsv")
	if !file.FileExists(addressFn) {
		return validate.Usage("Cannot find address file {0}", addressFn)
	}

	type Counters struct {
		nRead      int
		nProcessed int
	}
	counter := Counters{}
	skipped := []types.Grant{}

	// Define where to find addresses file
	grantReader, err := tokenomics.ReadGrants(addressFn)
	if err != nil {
		log.Fatal(err)
	}

	for {
		grant, err := grantReader.Read()
		if err == io.EOF {
			break
		}

		counter.nRead++
		if err != nil {
			log.Fatal(err)
		}
		if !grant.IsValid {
			skipped = append(skipped, grant)
			continue
		}

		counter.nProcessed++
		logger.Log(logger.Info, fmt.Sprintf("Updated data for %s", grant.Address))

		for _, chain := range chains {
			// Get some data we're going to need. Current state of the chain...
			meta := rpcClient.GetMetaData(chain, false)

			chainData := types.Chain{ChainName: chain}
			chainData.Counts, err = LineCounts(folder, chain, format, grant.Address)
			if err != nil {
				log.Fatal(err)
			}

			mon := monitor.NewMonitor(chain, grant.Address, false)
			chainData.FileSize = file.FileSize(mon.Path())
			if file.FileExists(mon.Path()) {
				err = mon.ReadHeader()
				if err != nil {
					log.Fatal(err)
				}

				if chainData.Counts.Appearances > 0 {
					apps := make([]index.AppearanceRecord, mon.Count())
					err = mon.ReadAppearances(&apps)
					if err != nil {
						log.Fatal(err)
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

					chainData.Balances = append(chainData.Balances, types.Balance{
						Asset:   "ETH",
						Balance: GetBalanceInEth(chain, grant.Address, meta.Latest),
					})
				}
			}

			mon.Close()
			chainData.Types = chainData.Counts.Types()
			grant.Chains = append(grant.Chains, chainData)
			grantMap[grant.Key] = grant
		}
	}

	sorted := []types.Grant{}
	for _, grant := range grantMap {
		sorted = append(sorted, grant)
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Address == sorted[j].Address {
			return sorted[i].GrantId < sorted[j].GrantId
		}
		return sorted[i].Address < sorted[j].Address
	})

	first := true
	fmt.Println("[")
	for i, grant := range sorted {
		if len(strings.Trim(grant.GrantId, " ")) == 0 {
			grant.GrantId = fmt.Sprintf("%04d", i)
		}
		if !first {
			fmt.Println(",")
		}
		str := grant.String()
		fmt.Println(str)
		first = false
	}
	fmt.Println("]")

	logger.Log(logger.Info, "nRead:", counter.nRead, "nProcessed:", counter.nProcessed)
	for sk := range skipped {
		logger.Log(logger.Info, "\t", sk)
	}

	return nil
}

func LineCounts(folder, chain, format, addr string) (types.Counts, error) {
	if !strings.HasSuffix(folder, "/") {
		folder += "/"
	}

	fileName := addr + "." + format

	base := path.Join(folder, "exports", chain)
	if !file.FolderExists(base) {
		return types.Counts{}, fmt.Errorf("data folder (%s) not found", base)
	}

	counts := types.Counts{}
	counts.Appearances, _ = file.LineCount(path.Join(base, "apps", fileName), true)
	if counts.Appearances > 0 {
		counts.Neighbors, _ = file.LineCount(folder+"exports/"+chain+"/neighbors/"+addr+"."+format, true)
		counts.Logs, _ = file.LineCount(folder+"exports/"+chain+"/logs/"+addr+"."+format, true)
		counts.Txs, _ = file.LineCount(folder+"exports/"+chain+"/txs/"+addr+"."+format, true)
		counts.Statements, _ = file.LineCount(folder+"exports/"+chain+"/statements/"+addr+"."+format, true)
	}
	return counts, nil
}

// TODO: BOGUS move to package
func ethFromWei(in big.Int) float64 {
	inF := new(big.Float).SetInt(&in)
	powI := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	powF := new(big.Float).SetInt(powI)
	out := inF.Quo(inF, powF)
	f, _ := out.Float64()
	return f
}

// TODO: BOGUS this should be generalized to the client itself instead of hidden in balanceClient
// balanceClient caches the client so we can call it many times without re-dialing it every time
var balanceClient = make(map[string]*ethclient.Client)
var clientLoaded = make(map[string]bool)

// GetBalanceInEth returns the balance of the given address at the given block
// TODO: BOGUS blockNum is ignored
// TODO: BOGUS what to do if we're running against a non-archive node?
func GetBalanceInEth(chain, address string, blockNum uint64) float64 {
	if !clientLoaded[chain] {
		provider := config.GetRpcProvider(chain)
		balanceClient[chain] = rpcClient.GetClient(provider)
		clientLoaded[chain] = true
	}
	val, _ := balanceClient[chain].BalanceAt(context.Background(), common.HexToAddress(address), big.NewInt(int64(blockNum)))
	if val == nil {
		return 0.0
	}
	return ethFromWei(*val)
}

func readExistingGrants(folder string) (map[string]types.Grant, error) {
	theMap := map[string]types.Grant{}

	path := path.Join(folder, "ui/src/theData.json")
	bytes, _ := ioutil.ReadFile(path)

	grants := []types.Grant{}
	_ = json.Unmarshal([]byte(bytes), &grants)

	for _, grant := range grants {
		grant.Key = grant.Address + "_" + grant.GrantId
		theMap[grant.Key] = grant
	}

	return theMap, nil
}
