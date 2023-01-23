package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	readers, err := loadReaders("testdata/readers.json")
	if err != nil {
		panic(err)
	}

	matchingBooks := findMatchingBooks(readers)

	fmt.Println(matchingBooks)
}

// loadReaders reads the file and returns the list of readers, and their beloved books, found therein.
func loadReaders(filePath string) ([]Reader, error) {
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

// findMatchingBooks returns books that are on more than one reader's shelf.
func findMatchingBooks(readers []Reader) []Book {
	booksOnShelves := make(map[Book]uint)

	// Register all books on shelves.
	for _, reader := range readers {
		for _, book := range reader.Books {
			booksOnShelves[book]++
		}
	}

	matchingBooks := make([]Book, 0)

	// Find books that were added to shelves more than once.
	for book, count := range booksOnShelves {
		if count > 1 {
			matchingBooks = append(matchingBooks, book)
		}
	}

	return matchingBooks
}