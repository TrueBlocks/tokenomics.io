package tokenomics

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

func ReadAddressFile(path string) (result []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	// Here we extract
	for _, record := range records {
		if !validate.IsValidAddress(record[0]) {
			err = fmt.Errorf("not a valid address: %s", record[0])
			return
		}

		result = append(result, record[0])
	}

	// Return a channel / put AddressRecord into a channel
	return
}
