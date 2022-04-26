package cmd

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
	"github.com/spf13/cobra"
)

func runCompress(cmd *cobra.Command, args []string) error {
	folder, chain, format := getOptions()

	addressFn := path.Join(folder, "./addresses.txt")
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

		log.Println("Compressing", grant.Address)

		tarFile, err := os.OpenFile(path.Join(folder, "./exports", chain, "zips", grant.Address+".tar"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		archive := tar.NewWriter(tarFile)
		defer func() {
			if err := archive.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		for _, dataType := range dataTypes {
			log.Println("\t", dataType)
			input, err := os.ReadFile(path.Join(folder, "./exports", chain, dataType, grant.Address+"."+format))
			if err != nil {
				log.Println(err)
			} else {
				archive.Write(input)
			}
		}

		log.Println("Done", grant.Address)
		tarFile.Close()
	}
	return nil
}
