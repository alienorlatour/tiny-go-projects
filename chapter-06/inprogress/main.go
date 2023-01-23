package main

import (
	"fmt"
)

func main() {
	readers, err := loadReaders("testdata/readers.json")
	if err != nil {
		panic(err)
	}

	matchingBooks := findMatchingBooks(readers)

	fmt.Println(matchingBooks)
}
