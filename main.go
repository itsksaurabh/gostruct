package main

import (
	"flag"
	"fmt"
	"itsksaurabh/getstruct/parser"
	"log"
	"os"
	"strings"
)

var (
	// file path to read from. Defaults to "./example.go"
	filePath = flag.String("fpath", "./example.go", "file path to read from")
)

func validateFilePath(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	flag.Parse()

	if !validateFilePath(*filePath) {
		panic("file does not exist")
	}

	names, err := parser.GetStructNames(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s List of Struct names in %s file %s\n", strings.Repeat("=", 25), *filePath, strings.Repeat("=", 25))
	for _, name := range names {
		fmt.Printf("* %s\n", name)
	}
}
