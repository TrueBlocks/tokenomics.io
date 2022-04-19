package tokenomics

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

var requiredColumns = []string{
	"address",
	"name",
}

type GrantReader struct {
	file      *os.File
	header    map[string]int
	csvReader csv.Reader
}

func (gr *GrantReader) Read() (types.Grant, error) {
	record, err := gr.csvReader.Read()
	if err == io.EOF {
		gr.file.Close()
	}
	if err != nil {
		return types.Grant{}, err
	}
	if !validate.IsValidAddress(record[0]) {
		err = fmt.Errorf("not a valid address: %s", record[0])
		return types.Grant{}, err
	}

	return types.Grant{
		Address: record[gr.header["address"]],
	}, nil
}

func ReadGrants(path string) (grantReader GrantReader, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	// read header
	headerRow, err := reader.Read()
	if err != nil {
		return
	}
	header := map[string]int{}
	for index, columnName := range headerRow {
		header[columnName] = index
	}
	// make sure the header is correct
	for _, required := range requiredColumns {
		_, ok := header[required]
		if !ok {
			err = fmt.Errorf(`required column "%s" missing in file %s`, required, path)
			return
		}
	}

	return GrantReader{
		file:      file,
		header:    header,
		csvReader: *reader,
	}, nil
}
