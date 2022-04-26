package internal

import (
	"io"
	"log"
	"os"
	"path"
	"strings"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
	"github.com/spf13/cobra"
)

func RunCompress(cmd *cobra.Command, args []string) error {
	folder, chain, format := getOptions(cmd.Parent())
	zipFolder := path.Join(folder, "./exports", chain, "zips")

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

		// We create a temporary folder (named after the address) and copy the files over
		// there so we get the folder structure we want when the .tar.gz file is extracted
		// We will delete this in a moment
		tmpFolder := path.Join(zipFolder, grant.Address)
		file.EstablishFolder(tmpFolder)

		files := []string{}
		totalSize := int64(0)
		for _, dataType := range dataTypes {
			sourceFn := path.Join(folder, "./exports", chain, dataType, grant.Address+".") + format
			thisSize := file.FileSize(sourceFn)
			if thisSize > 0 {
				totalSize += thisSize
				destFn := tmpFolder + "/" + dataType + "." + format
				file.Copy(destFn, sourceFn)
				files = append(files, destFn)
			}
		}

		if len(files) == 0 {
			// logger.Log(logger.Info, "Nothing to archive for", grant.Address)
			continue
		}

		tarFn := path.Join(folder, "./exports", chain, "zips", grant.Address+".tar.gz")
		out, err := os.OpenFile(tarFn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalln("Error writing archive:", err)
		}

		// Create the archive and write the output to the "out" Writer
		err = file.CreateArchive(files, out, true, zipFolder)
		if err != nil {
			log.Fatalln("Error creating archive:", err)
		}
		theFolder := strings.Replace(tarFn, ".tar.gz", "", -1)
		os.Remove(theFolder) // remove the now empty folder
		out.Close()          // do not defer, we want to close it now

		logger.Log(logger.Info, "Compressed", len(files), "files for address", grant.Address, "(size:", totalSize, ")")
	}

	return nil
}
