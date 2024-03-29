package types

import (
	"encoding/json"
	"strings"
)

type Counts struct {
	Appearances int `json:"appearanceCount"`
	Logs        int `json:"logCount"`
	Txs         int `json:"transactionCount"`
	Neighbors   int `json:"neighborCount"`
	Statements  int `json:"statementsCount"`
}

func (c Counts) String() string {
	ret, _ := json.MarshalIndent(c, "", "  ")
	return string(ret)
}

func (c Counts) Types() string {
	ret := ""
	if c.Appearances > 0 {
		ret += "apps,"
	}
	if c.Logs > 0 {
		ret += "logs,"
	}
	if c.Txs > 0 {
		ret += "txs,"
	}
	if c.Neighbors > 0 {
		ret += "neighbors,"
	}
	if c.Statements > 0 {
		ret += "statements,"
	}
	return strings.Trim(ret, ",")
}

type Appearance struct {
	Bn        int    `json:"bn,omitempty"`
	TxId      int    `json:"txId,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
	Date      string `json:"date,omitempty"`
}

type Balance struct {
	Asset   string  `json:"asset"`
	Balance float64 `json:"balance"`
}

type Chain struct {
	ChainName  string     `json:"chainName,omitempty"`
	FirstApp   Appearance `json:"firstAppearance,omitempty"`
	LatestApp  Appearance `json:"latestAppearance,omitempty"`
	FileSize   int64      `json:"fileSize,omitempty"`
	BlockRange int        `json:"blockRange,omitempty"`
	Counts     Counts     `json:"counts,omitempty"`
	Balances   []Balance  `json:"balances,omitempty"`
	Types      string     `json:"types,omitempty"`
}

func (c Chain) HasRecords() bool {
	return (c.Counts.Appearances + c.Counts.Logs + c.Counts.Txs + c.Counts.Neighbors + c.Counts.Statements) > 0
}

type Grant struct {
	GrantId     string  `json:"grantId,omitempty"`
	Address     string  `json:"address,omitempty"`
	Name        string  `json:"name,omitempty"`
	Tag         string  `json:"tag,omitempty"`
	LastUpdated int64   `json:"lastUpdated,omitempty"`
	IsActive    bool    `json:"isActive,omitempty"`
	IsCore      bool    `json:"core"`
	IsValid     bool    `json:"isValid,omitempty"`
	Chains      []Chain `json:"chainData,omitempty"`
	Key         string  `json:"-"`
}

func (g Grant) String() string {
	ret, _ := json.MarshalIndent(g, "", "  ")
	return string(ret)
}
