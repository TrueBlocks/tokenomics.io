// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package exportPkg

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TrueBlocks/tokenomics.io/gitcoin/backend/pkg/rpcClient"
)

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
