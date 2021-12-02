// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package rpcClient

import (
	"context"
	"math/big"
)

func ethFromWei(in big.Int) float64 {
	inF := new(big.Float).SetInt(&in)
	powI := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	powF := new(big.Float).SetInt(powI)
	out := inF.Quo(inF, powF)
	f, _ := out.Float64()
	return f
}

func GetBalanceInEth(address string, bn uint64) float64 {
	client := Get()
	val, _ := client.BalanceAt(context.Background(), HexToAddress(address), nil)
	if val == nil {
		return 0.0
	}
	return ethFromWei(*val)
}
