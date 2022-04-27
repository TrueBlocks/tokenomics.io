package internal

import (
	"fmt"
	"io"
	"log"
	"path"
	"strings"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
	"github.com/spf13/cobra"
)

func RunNames(cmd *cobra.Command, args []string) error {
	folder, _, _ := getOptions(cmd.Parent())

	tag, err := cmd.Flags().GetString("tag")
	if err != nil {
		log.Fatal(err)
	}
	if len(tag) == 0 {
		tag = "30-Monitored"
	}

	addressFn := path.Join(folder, "./addresses.tsv")
	if !file.FileExists(addressFn) {
		return validate.Usage("Cannot find address file {0}", addressFn)
	}

	// Define where to find addresses file
	grantReader, err := tokenomics.ReadGrants(addressFn)
	if err != nil {
		log.Fatal(err)
	}

	for {
		grant, err := grantReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !grant.IsValid {
			continue
		}

		str := "addName \"{ADDRESS}\" \"{NAME}\" \"{TAG}\" \"{SOURCE}\""
		str = strings.Replace(str, "{ADDRESS}", strings.ToLower(grant.Address), -1)
		str = strings.Replace(str, "{NAME}", grant.Name, -1)
		str = strings.Replace(str, "{TAG}", tag, -1)
		str = strings.Replace(str, "{SOURCE}", "Tokenomics", -1)
		fmt.Println(str)
	}
	return nil
}
