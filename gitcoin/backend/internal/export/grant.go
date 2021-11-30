// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package exportPkg

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/rpcClient"
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

type GrantSummary struct {
	Key           uint64   `json:"key"`
	Date          string   `json:"date"`
	LastBlock     uint64   `json:"last_block"`
	LastTs        uint64   `json:"last_ts"`
	Type          string   `json:"type"`
	Id            uint64   `json:"grant_id"`
	Address       string   `json:"address"`
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	Logo          string   `json:"logo"`
	TxCount       uint64   `json:"tx_cnt"`
	LogCount      uint64   `json:"log_cnt"`
	DonationCount uint64   `json:"donation_cnt"`
	Matched       float64  `json:"matched"`
	Claimed       float64  `json:"claimed"`
	Balances      Balances `json:"balances"`
	IsCore        bool     `json:"core"`
}
type GrantSummaries []GrantSummary

func (g *GrantSummary) ToJSON() string {
	str, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(str)
}

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

func (s *GrantSummary) FromGrant(grant *Grant) {
	s.Key = 0
	s.Date = ""
	s.LastBlock = 0
	s.LastTs = 0
	s.Type = "logs"
	s.Id = grant.Id
	s.Address = strings.ToLower(grant.AdminAddress)
	s.Name = properTitle(grant.Title)
	id := fmt.Sprintf("%d", grant.Id)
	s.Slug = "https://gitcoin.co/grants/" + id + "/" + grant.Slug
	s.Logo = ""
	if grant.Logo != nil {
		s.Logo = *grant.Logo
	}
	s.TxCount = 0
	s.LogCount = 0
	s.DonationCount = 0
	s.Matched = 0
	s.Claimed = 0
	s.Balances = []Balance{Balance{Asset: "ETH", Balance: 1.0}}
	s.IsCore = false
}

type Appearance struct {
	Bn   uint32 `json:"bn"`
	TxId uint32 `json:"tx_id"`
}

type Monitor struct {
	Address    string     `json:"address"`
	First      Appearance `json:"firstAppearance"`
	Latest     Appearance `json:"latestAppearance"`
	LastUpdate uint32     `json:"lastUpdated"`
	Age        uint32     `json:"ageInBlocks"`
	Range      uint32     `json:"blockRange"`
	Size       int64      `json:"fileSize"`
	Count      int64      `json:"appearanceCount"`
	Neighbors  int64      `json:"neighborCount"`
}

func GetMonitorStats(grantId string, address string) (*Monitor, error) {
	monitor := &Monitor{Address: address, First: Appearance{Bn: 10, TxId: 20}, Latest: Appearance{Bn: 30, TxId: 40}, Size: 100, Count: 222}
	monitorPath := fmt.Sprintf("/Users/jrush/Library/Application Support/TrueBlocks/cache/monitors/%s.acct.bin", address)
	fileStat, err := os.Stat(monitorPath)
	if err != nil {
		return nil, err
	}
	monitor.Size = fileStat.Size()
	monitor.Count = fileStat.Size() / 8
	monitorFile, err := os.Open(monitorPath)
	if err != nil {
		return nil, err
	}
	defer monitorFile.Close()

	err = binary.Read(monitorFile, binary.LittleEndian, &monitor.First.Bn)
	if err != nil {
		return nil, err
	}
	err = binary.Read(monitorFile, binary.LittleEndian, &monitor.First.TxId)
	if err != nil {
		return nil, err
	}
	monitorFile.Seek(-8, 2)
	err = binary.Read(monitorFile, binary.LittleEndian, &monitor.Latest.Bn)
	if err != nil {
		return nil, err
	}
	err = binary.Read(monitorFile, binary.LittleEndian, &monitor.Latest.TxId)
	if err != nil {
		return nil, err
	}
	monitor.Range = monitor.Latest.Bn - monitor.First.Bn + 1
	var meta rpcClient.Meta
	monitor.Age = uint32(meta.Latest()) - monitor.First.Bn

	return monitor, nil
}

func (m *Monitor) ToJSON() string {
	str, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(str)
}
