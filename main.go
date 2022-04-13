package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var folder string
	flag.StringVar(&folder, "folder", "", "which folder to process")
	flag.Parse()

	if folder == "" {
		fmt.Println("Please provide a folder name to process")
		os.Exit(1)
	}
}
