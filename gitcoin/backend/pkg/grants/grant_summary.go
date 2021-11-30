package grants

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
