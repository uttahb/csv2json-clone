package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type inputFile struct {
	filePath  string
	separator string
	pretty    bool
}

func getFileData() (inputFile, error) {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("A file path argument is required")
	}
	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Getting pretty json")
	flag.Parse()
	fmt.Println(flag.Args())
	fileLocation := flag.Arg(0)
	if !(*separator == "comma" || *separator == "semicolumn") {
		return inputFile{}, errors.New("Only comma or semicolon separators are allowed")
	}
	return inputFile{fileLocation, *separator, *pretty}, nil
}
func checkIfValidFile(filename string) (bool, error) {
	if fileExtension := filepath.Ext(filename); fileExtension != "csv" {
		return false, fmt.Errorf("file %s is not csv", filename)
	}
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %s does not exist", filename)
	}
	return true, nil
}
func main() {
	getFileData()
}
