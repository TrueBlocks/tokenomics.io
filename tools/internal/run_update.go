package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"path"
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
	folder, chain, format := getOptions(cmd.Parent())

	addressFn := path.Join(folder, "./addresses.txt")
	if !file.FileExists(addressFn) {
		return validate.Usage("Cannot find address file {0}", addressFn)
	}

	// Define where to find addresses file
	grantReader, err := tokenomics.ReadGrants(addressFn)
	if err != nil {
		log.Fatal(err)
	}

	meta := rpcClient.GetMetaData("mainnet", false)
	grants := []types.Grant{}
	for {
		grant, err := grantReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !grant.IsValid {
			continue
		}

		logger.Log(logger.Info, fmt.Sprintf("Updated data for %s", grant.Address))

		chainData := types.Chain{ChainName: "mainnet"}
		chainData.Counts, err = LineCounts(folder, chain, format, grant.Address)
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
			if chainData.Counts.Appearances > 0 {
				apps := make([]index.AppearanceRecord, mon.Count())
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
				chainData.Balances = append(chainData.Balances, types.Balance{
					Asset:   "ETH",
					Balance: GetBalanceInEth("mainnet", grant.Address, meta.Latest),
				})
			}
		}
		mon.Close()
		chainData.Types = chainData.Counts.Types()
		grant.Chains = append(grant.Chains, chainData)

		grants = append(grants, grant)
	}

	fmt.Println("[")
	first := true
	for _, grant := range grants {
		if grant.Chains[0].HasRecords() {
			if !first {
				fmt.Println(",")
			}
			str := grant.String()
			fmt.Println(str)
			first = false
		}
	}
	fmt.Println("]")
	return nil
}

func LineCounts(folder, chain, format, addr string) (types.Counts, error) {
	if !strings.HasSuffix(folder, "/") {
		folder += "/"
	}
	base := "./" + folder + "exports/" + chain
	if !file.FolderExists(base) {
		fmt.Println("SOMSOMTEOMTOME")
		return types.Counts{}, fmt.Errorf("data folder (%s) not found", base)
	}
	counts := types.Counts{}
	counts.Appearances, _ = file.LineCount(folder+"exports/"+chain+"/apps/"+addr+"."+format, true)
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
var balanceClient *ethclient.Client
var clientLoaded = false

// GetBalanceInEth returns the balance of the given address at the given block
// TODO: BOGUS blockNum is ignored
// TODO: BOGUS what to do if we're running against a non-archive node?
func GetBalanceInEth(chain, address string, blockNum uint64) float64 {
	if !clientLoaded {
		provider := config.GetRpcProvider(chain)
		balanceClient = rpcClient.GetClient(provider)
		clientLoaded = true
	}
	val, _ := balanceClient.BalanceAt(context.Background(), common.HexToAddress(address), big.NewInt(int64(blockNum)))
	if val == nil {
		return 0.0
	}
	return ethFromWei(*val)
}
