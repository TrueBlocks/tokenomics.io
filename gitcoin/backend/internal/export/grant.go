// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package exportPkg

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	tslibPkg "github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/tslib"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
)

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

func properTitle(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

type Appearance struct {
	Bn   uint32 `json:"bn"`
	TxId uint32 `json:"txId"`
	Ts   uint64 `json:"timestamp"`
}

type Monitor struct {
	Id         uint64     `json:"grantId"`
	Address    string     `json:"address"`
	Name       string     `json:"name"`
	Slug       string     `json:"slug"`
	First      Appearance `json:"firstAppearance"`
	Latest     Appearance `json:"latestAppearance"`
	LastUpdate uint64     `json:"lastUpdated"`
	// Age         uint32     `json:"ageInBlocks"`
	Range       uint32   `json:"blockRange"`
	Size        int64    `json:"fileSize"`
	Count       int64    `json:"appearanceCount"`
	Neighbors   int64    `json:"neighborCount"`
	Types       string   `json:"types"`
	LogCount    int64    `json:"logCount"`
	DonationCnt int64    `json:"donationCount"`
	Matched     float64  `json:"matched"`
	Claimed     float64  `json:"claimed"`
	Balances    Balances `json:"balances,omitempty"`
	Core        bool     `json:"core"`
}

func (m *Monitor) ReadRangeAndAge(monitorPath string) error {
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

	m.Range = m.Latest.Bn - m.First.Bn + 1
	// var meta rpcClient.Meta
	// m.Age = uint32(meta.Latest()) - m.First.Bn
	return nil
}

const path = "/Users/jrush/Library/Application Support/TrueBlocks/cache/monitors/"

func GetMonitorStats(grantId string, grant *Grant) (*Monitor, error) {
	monitor := &Monitor{
		Address: grant.AdminAddress,
		First: Appearance{
			Bn:   10,
			TxId: 20,
		},
		Latest: Appearance{
			Bn:   30,
			TxId: 40,
		},
		Size:  100,
		Count: 222,
	}

	monitorPath := fmt.Sprintf(path+"%s.acct.bin", grant.AdminAddress)
	if !file.FileExists(monitorPath) {
		return nil, errors.New("file not exist")
	}

	fileStat, err := os.Stat(monitorPath)
	if err != nil {
		return nil, err
	}

	monitor.Slug = "https://gitcoin.co/grants/" + fmt.Sprintf("%d", grant.Id) + "/" + grant.Slug
	monitor.Id = grant.Id
	monitor.Core = false
	monitor.Name = strings.Replace(grant.Title, "'", "", -1)
	monitor.Types = "txs,logs,neighbors"
	monitor.Size = fileStat.Size()
	monitor.Count = fileStat.Size() / 8

	err = monitor.ReadRangeAndAge(monitorPath)
	if err != nil {
		return nil, err
	}

	lastBlockPath := path + grant.AdminAddress + ".last.txt"
	file, err := os.Open(lastBlockPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	line, err := r.ReadBytes('\n')
	str := strings.Replace(string(line), "\n", "", -1)
	if err != nil {
		return nil, err
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}
	monitor.LastUpdate = uint64(val)

	return monitor, nil
}

func (m *Monitor) ToJSON() string {
	str, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(str)
}
