package tokenomics

import (
	"io"
	"testing"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/types"
)

func TestReadAddressFile(t *testing.T) {
	reader, err := ReadGrants("./testdata/addresses.tsv")
	if err != nil {
		t.Fatal(err)
	}

	firstGrant := types.Grant{}
	readCount := 0
	for {
		grant, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal("error while reading records", err)
		}

		if firstGrant.Address == "" {
			firstGrant = grant
		}
		readCount++
	}

	if readCount != 21 {
		t.Fatalf("expected len to be 50, but got %d", readCount)
	}

	if firstGrant.Address != "0x001bae3291db1f401771a1eaa8f0078534177bcc" {
		t.Fatal("wrong address read", firstGrant.Address)
	}
}

func TestReadAddressFile_Invalid(t *testing.T) {
	reader, err := ReadGrants("./testdata/addresses_invalid.tsv")
	if err != nil {
		t.Fatal(err)
	}

	_, err = reader.Read()
	if err == nil {
		t.Fatal("expected error")
	}
}
