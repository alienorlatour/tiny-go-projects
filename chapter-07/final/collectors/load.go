package collectors

import (
	"bufio"
	"encoding/json"
	"os"
)

// Load reads the file and returns the list of collectors, and their beloved books, found therein.
func Load[T comparable](filePath string) (Collectors[T], error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Use a buffered reader to control system calls.
	bf := bufio.NewReaderSize(f, 1024)

	// Declare the variable in which the file will be decoded.
	var colls Collectors[T]

	// Decode the file and store the content in the variable colls.
	err = json.NewDecoder(bf).Decode(&colls)
	if err != nil {
		return nil, err
	}

	return colls, nil
}
