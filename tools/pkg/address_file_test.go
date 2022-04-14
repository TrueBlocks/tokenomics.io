package tokenomics

import (
	"testing"
)

func TestReadAddressFile(t *testing.T) {
	result, err := ReadAddressFile("./testdata/addresses.csv")
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 50 {
		t.Fatalf("expected len to be 50, but got %d", len(result))
	}
}

func TestReadAddressFile_Invalid(t *testing.T) {
	_, err := ReadAddressFile("./testdata/addresses_invalid.csv")
	if err == nil {
		t.Fatal("expected error")
	}
}
