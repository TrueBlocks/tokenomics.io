package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk("../", visitGrant)
}

func visitGrant(path string, info os.FileInfo, err error) error {
	if strings.Contains(path, "code") {
		return nil

	} else if strings.Contains(path, ".json") {

		jsonFile, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer jsonFile.Close()

		var grants Grants
		byteValue, _ := ioutil.ReadAll(jsonFile)
		err = json.Unmarshal(byteValue, &grants)
		if err != nil {
			fmt.Println(err)
			return err
		}

		outPath := "./two/" + path[3:]
		fmt.Fprintf(os.Stderr, "Processing: %s\n", outPath)
		f, err := os.Create(outPath)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		if len(grants) > 0 {
			res := fmt.Sprintf("[%s]", grants[0].ToJSON())
			_, err = w.WriteString(res)
		} else {
			_, err = w.WriteString("[]")
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		w.Flush()
	}
	return nil
}
