package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	readers, err := loadData("testdata/readers.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(readers)
}

// loadData reads the file and returns the list of readers, and their beloved books, found therein.
func loadData(filePath string) ([]Reader, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	readers := make([]Reader, 0)
	err = json.Unmarshal(contents, &readers)
	if err != nil {
		return nil, err
	}

	return readers, nil
}
