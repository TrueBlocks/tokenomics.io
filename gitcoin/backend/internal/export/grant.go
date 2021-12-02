// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package exportPkg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/rpcClient"
	tslibPkg "github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/tslib"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
)

const pathToMonitors = "/Users/jrush/Library/Application Support/TrueBlocks/cache/monitors/"

// Grant is one of the Gitcoin Grants
type Grant struct {
	Id                            uint64   `json:"id"`
	Active                        bool     `json:"active"`
	Title                         string   `json:"title"`
	Slug                          string   `json:"slug"`
	Description                   string   `json:"description"`
	ReferenceUrl                  string   `json:"reference_url"`
	Logo                          *string  `json:"logo"`
	AdminAddress                  string   `json:"admin_address"`
	AmountReceived                string   `json:"amount_received"`
	TokenAddress                  string   `json:"token_address"`
	TokenSymbol                   string   `json:"token_symbol"`
	ContractAddress               string   `json:"contract_address"`
	Meta                          MetaData `json:"metadata"`
	Network                       string   `json:"network"`
	RequiredGasPrice              string   `json:"required_gas_price"`
	AdminProfile                  Profile  `json:"admin_profile"`
	TeamMembers                   Profiles `json:"team_members"`
	ClrPredictionCurve            Points3d `json:"clr_prediction_curve"`
	ClrRoundNum                   string   `json:"clr_round_num"`
	IsClrActive                   bool     `json:"is_clr_active"`
	AmountReceivedInRound         string   `json:"amount_received_in_round"`
	PositiveRoundContributorCount uint64   `json:"positive_round_contributor_count"`
	Monitor                       Monitor  `json:"monitor"`
}
type Grants []Grant

func (g *Grant) ToJSON() string {
	str, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(str)
}

// Points are 2d and 3d arrays of floats
type Point2d [2]float64
type Points2d []Point2d
type Point3d [3]float64
type Points3d []Point3d

// MetaData carries information about the grant
type MetaData struct {
	Gem                                int64            `json:"gem"`
	Related                            *Points2d        `json:"related,omitempty"`
	Upcoming                           int64            `json:"upcoming"`
	WallOfLove                         [][2]interface{} `json:"wall_of_love"`
	UnsubscribedProfiles               []uint64         `json:"unsubscribed_profiles,omitempty"`
	LastCalcTimeRelated                *float64         `json:"last_calc_time_related,omitempty"`
	LastCalcTimeContributorCounts      float64          `json:"last_calc_time_contributor_counts"`
	LastCalcTimeSybilAndContribAmounts float64          `json:"last_calc_time_sybil_and_contrib_amounts"`
}

// Profile is a Github user
type Profile struct {
	Id           uint64            `json:"id"`
	Url          string            `json:"url"`
	Name         string            `json:"name"`
	Handle       string            `json:"handle"`
	Keywords     []string          `json:"keywords"`
	Position     uint64            `json:"position"`
	AvatarUrl    string            `json:"avatar_url"`
	GithubUrl    string            `json:"github_url"`
	TotalEarned  float64           `json:"total_earned"`
	Oranizations map[string]uint64 `json:"organizations"`
}
type Profiles []Profile

// Love is donor comments on a Grant
type Love map[string]uint64
type Loves []Love

func (g *Grant) GetGrant(grantId string) error {

	jsonFile, err := os.Open(grantId)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)

	// data on disc is stored as an array, but only contains a single grant
	var tmpArray Grants
	err = json.Unmarshal(byteData, &tmpArray)
	if err != nil {
		return err

	} else {
		if len(tmpArray) > 0 {
			*g = tmpArray[0]
			g.AdminAddress = strings.ToLower(g.AdminAddress)
		}
	}

	// It's okay for the grant to be non-existant
	return nil
}

type Balance struct {
	Asset   string  `json:"asset"`
	Balance float64 `json:"balance"`
}
type Balances []Balance

// func properTitle(input string) string {
// 	words := strings.Split(input, " ")
// 	smallwords := " a an on the to "

// 	for index, word := range words {
// 		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
// 			words[index] = word
// 		} else {
// 			words[index] = strings.Title(word)
// 		}
// 	}
// 	return strings.Join(words, " ")
// }

type Appearance struct {
	Bn      uint32 `json:"bn"`
	TxId    uint32 `json:"txId"`
	Ts      uint64 `json:"timestamp"`
	DateStr string `json:"date",omitempty`
}

type Monitor struct {
	Id            uint64     `json:"grantId"`
	Address       string     `json:"address"`
	Name          string     `json:"name"`
	Slug          string     `json:"slug"`
	First         Appearance `json:"firstAppearance"`
	Latest        Appearance `json:"latestAppearance"`
	LastUpdate    uint64     `json:"lastUpdated"`
	Range         uint32     `json:"blockRange"`
	Size          int64      `json:"fileSize"`
	Count         int64      `json:"appearanceCount"`
	NeighborCount int64      `json:"neighborCount"`
	Types         string     `json:"types"`
	LogCount      int64      `json:"logCount"`
	Matched       float64    `json:"matched"`
	Claimed       float64    `json:"claimed"`
	Balances      Balances   `json:"balances,omitempty"`
	Core          bool       `json:"core"`
}

func CountLines(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	cnt, err := LineCounter(file)
	if err != nil {
		return 0, err
	}
	return int64(cnt), nil
}

func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

func (m *Monitor) ReadRange(monitorPath string) error {
	monitorFile, err := os.Open(monitorPath)
	if err != nil {
		return err
	}
	defer monitorFile.Close()

	err = binary.Read(monitorFile, binary.LittleEndian, &m.First.Bn)
	if err != nil {
		return err
	}
	err = binary.Read(monitorFile, binary.LittleEndian, &m.First.TxId)
	if err != nil {
		return err
	}
	m.First.Ts, _ = tslibPkg.TsFromBn(uint64(m.First.Bn))
	m.First.DateStr, _ = tslibPkg.DateFromTs(m.First.Ts)
	m.First.DateStr = strings.Replace(m.First.DateStr, "T", " ", -1)

	monitorFile.Seek(-8, 2)
	err = binary.Read(monitorFile, binary.LittleEndian, &m.Latest.Bn)
	if err != nil {
		return err
	}
	err = binary.Read(monitorFile, binary.LittleEndian, &m.Latest.TxId)
	if err != nil {
		return err
	}
	m.Latest.Ts, _ = tslibPkg.TsFromBn(uint64(m.Latest.Bn))
	m.Latest.DateStr, _ = tslibPkg.DateFromTs(m.Latest.Ts)
	m.Latest.DateStr = strings.Replace(m.Latest.DateStr, "T", " ", -1)

	m.Range = m.Latest.Bn - m.First.Bn + 1
	return nil
}

func (m *Monitor) getLastUpdate() (uint64, error) {
	lastBlockPath := pathToMonitors + m.Address + ".last.txt"
	file, err := os.Open(lastBlockPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	line, err := r.ReadBytes('\n')
	str := strings.Replace(string(line), "\n", "", -1)
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

func GetMonitorStats(grantId string, grant *Grant) (*Monitor, error) {
	monitor := &Monitor{Address: grant.AdminAddress}

	monitorPath := fmt.Sprintf(pathToMonitors+"%s.acct.bin", grant.AdminAddress)
	if !file.FileExists(monitorPath) {
		return nil, errors.New("file does not exist: " + monitorPath)
	}

	fileStat, err := os.Stat(monitorPath)
	if err != nil {
		return nil, err
	}

	monitor.Id = grant.Id
	if len(grant.Slug) > 0 {
		monitor.Slug = fmt.Sprintf("https://gitcoin.co/grants/%d/%s", grant.Id, grant.Slug)
	}
	monitor.Name = strings.Replace(grant.Title, "'", "", -1)
	monitor.Types = "txs,logs,neighbors"
	monitor.Size = fileStat.Size()
	monitor.Count = fileStat.Size() / 8
	monitor.Core = strings.Contains(grantId, "core")
	pp := pathToData + "neighbors/" + monitor.Address + ".csv"
	monitor.NeighborCount, _ = CountLines(pp)
	if monitor.NeighborCount > 0 {
		monitor.NeighborCount -= 1 // for header
	}
	pp = pathToData + "logs/" + monitor.Address + ".csv"
	monitor.LogCount, _ = CountLines(pp)
	if monitor.LogCount > 0 {
		monitor.LogCount -= 1 // for header
	}
	meta := rpcClient.GetMeta(false)
	bal := rpcClient.GetBalanceInEth(monitor.Address, meta.Latest())
	monitor.Balances = append(monitor.Balances, Balance{Asset: "ETH", Balance: bal})

	err = monitor.ReadRange(monitorPath)
	if err != nil {
		return nil, err
	}

	monitor.LastUpdate, err = monitor.getLastUpdate()
	if err != nil {
		return nil, err
	}

	return monitor, nil
}

func (m *Monitor) ToJSON() string {
	str, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(str)
}
