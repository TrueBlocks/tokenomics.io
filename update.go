package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/TrueBlocks/tokenomics.io/pkg/file"
	"github.com/TrueBlocks/tokenomics.io/pkg/types"
)

func main() {
	folder := "wallets"
	chain := "mainnet"

	grants := []types.Grant{}

	lines := file.AsciiFileToLines(folder + "/addresses.txt")
	for i, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) < 2 {
			parts = append(parts, fmt.Sprintf("%04d", i))
		}
		if len(parts) < 3 {
			parts = append(parts, "Name "+parts[1])
		}
		if len(parts) < 4 {
			parts = append(parts, "Active")
		}
		if len(parts) < 5 {
			parts = append(parts, "false")
		}
		if len(parts) < 6 {
			parts = append(parts, strings.ToLower(strings.Replace(parts[2], " ", "-", -1)))
		}

		grant := types.Grant{
			GrantId:     parts[1],
			Address:     parts[0],
			Name:        parts[2],
			IsActive:    parts[3] == "Active",
			Core:        parts[4] == "true",
			Slug:        parts[5],
			LastUpdated: time.Now().Unix(),
		}

		chainData := types.Chain{Name: "mainnet"}
		var err error
		chainData.Counts, err = file.LineCounts(folder, chain, grant.Address)
		if err != nil {
			log.Fatal(err)
		}
		chainData.Types = chainData.Counts.Types()
		grant.Chains = append(grant.Chains, chainData)

		grants = append(grants, grant)
	}

	for _, grant := range grants {
		if grant.Chains[0].HasRecords() {
			fmt.Println(grant)
		}
	}
}
