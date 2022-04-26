package tokenomics

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

var requiredColumns = []string{
	"address",
	"id",
	"name",
	"active",
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

	isActive := record[gr.header["active"]] == "true"
	isCore := record[gr.header["core"]] == "true"
	isValid := validate.IsValidAddress(record[0]) && !validate.IsZeroAddress(record[0])
	return types.Grant{
		Address:  strings.ToLower(record[gr.header["address"]]),
		Name:     record[gr.header["name"]],
		Slug:     record[gr.header["slug"]],
		IsActive: isActive,
		IsCore:   isCore,
		IsValid:  isValid,
	}, nil
}

func ReadGrants(path string) (GrantReader, error) {
	file, err := os.Open(path)
	if err != nil {
		return GrantReader{}, err
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	// read header
	headerRow, err := reader.Read()
	if err != nil {
		return GrantReader{}, err
	}
	header := map[string]int{}
	for index, columnName := range headerRow {
		header[columnName] = index
	}

	for _, required := range requiredColumns {
		_, ok := header[required]
		if !ok {
			err = fmt.Errorf(`required column "%s" missing in file %s`, required, path)
			return GrantReader{}, err
		}
	}

	grantReader := GrantReader{
		file:      file,
		header:    header,
		csvReader: *reader,
	}
	return grantReader, nil
}
