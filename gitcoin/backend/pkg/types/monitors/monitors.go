// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package monitorsPkg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	tslibPkg "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/tslib"
)

type Balance struct {
	Asset   string  `json:"asset"`
	Balance float64 `json:"balance"`
}

type Appearance struct {
	Bn      uint32 `json:"bn"`
	TxId    uint32 `json:"txId"`
	Ts      uint64 `json:"timestamp"`
	DateStr string `json:"date,omitempty"`
}

type Monitor struct {
	Id            uint64     `json:"grantId"`
	Active        bool       `json:"active"`
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
	Balances      []Balance  `json:"balances,omitempty"`
	Core          bool       `json:"core"`
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

	monitorFile.Seek(-8, io.SeekEnd)
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

	m.Range = m.Latest.Bn - m.First.Bn + 1
	return nil
}

func (m *Monitor) GetLastUpdate() (uint64, error) {
	pathToMonitors := config.ReadTrueBlocks().Settings.CachePath + "monitors/"
	lastBlockPath := pathToMonitors + m.Address + ".last.txt"
	file, err := os.Open(lastBlockPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	line, err := r.ReadBytes('\n')
	// str := strings.Replace(string(line), "\n", "", -1)
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(string(line))
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

func (m *Monitor) ToJSON() string {
	str, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(str)
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
