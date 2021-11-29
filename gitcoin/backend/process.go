package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var grants Grants

	root := "../data/raw"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".json") {
			jsonFile, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				return err
			}
			defer jsonFile.Close()

			fmt.Fprintf(os.Stderr, "Processing: %s\n", path)

			contents, _ := ioutil.ReadAll(jsonFile)

			var grantArray Grants // data is stored as an array, but only contains a single grant
			err = json.Unmarshal(contents, &grantArray)
			if err != nil {
				fmt.Println(err)
				return err
			}

			if len(grantArray) > 0 {
				grants = append(grants, grantArray[0])
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, grant := range grants {
		fmt.Printf("%s\n", strings.ToLower(grant.AdminAddress))
		for _, member := range grant.TeamMembers {
			str, err := json.Marshal(&member)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("%s\n", string(str))
			}
		}
		// fmt.Printf("%s: %s...\n", grant.AdminAddress, grant.ToJSON()[:40])
		// fmt.Println("../temp/" + path[12:])
		// fmt.Fprintf(os.Stderr, "Processing: %s\n", outPath)
		// f, err := os.Create(outPath)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		// defer f.Close()
		// w := bufio.NewWriter(f)
		// if len(grants) > 0 {
		// 	res := fmt.Sprintf("[%s]", grants[0].ToJSON())
		// 	_, err = w.WriteString(res)
		// } else {
		// 	_, err = w.WriteString("[]")
		// }
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		// w.Flush()
	}
}
