package internal

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	tokenomics "github.com/TrueBlocks/tokenomics.io/tools/pkg"
	"github.com/TrueBlocks/tokenomics.io/tools/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
	"github.com/spf13/cobra"
)

func compressCombined(folder string, chain string, format string) error {
	inputDir := path.Join(folder, chain, "combined")
	outputPath := path.Join(folder, chain, "zips/combined")
	err := file.EstablishFolder(outputPath)
	if err != nil {
		return err
	}

	err = filepath.WalkDir(inputDir, func(fp string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		if filepath.Ext(fp) != format {
			return nil
		}

		// inputPaths = append(inputPaths, path)
		name := filepath.Base(fp)
		outName := strings.Replace(name, format, ".tar.gz", 1)
		outFile, err := os.Create(path.Join(outputPath, outName))
		defer outFile.Close()
		if err != nil {
			return err
		}
		err = file.CreateArchive(
			[]string{fp},
			outFile,
			false,
			"",
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func RunCompress(cmd *cobra.Command, args []string) error {
	folder, chains, format := getOptions(cmd.Parent())

	addressFn := path.Join(folder, "./addresses.tsv")
	if !file.FileExists(addressFn) {
		return validate.Usage("Cannot find address file {0}", addressFn)
	}

	for _, chain := range chains {
		// Compress combined files first
		err := compressCombined(folder, chain, format)
		if err != nil {
			log.Fatal("Error while compressing combined files:", err)
		}

		// Define where to find addresses file
		gr, err := tokenomics.NewGrantReader(addressFn)
		if err != nil {
			log.Fatal(err)
		}

		for {
			grant, err := gr.Read()
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
			zipFolder := path.Join(folder, "./exports", chain, "zips")
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
	}

	return nil
}
