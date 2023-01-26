package main

import (
	"encoding/json"
	"io"
	"os"
	"sort"
)

// A Reader contains the list of books on a reader's shelf.
type Reader struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a reader's shelf.
type Book struct {
	Authors string `json:"authors"`
	Title   string `json:"title"`
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

	sort.Slice(matchingBooks, func(i, j int) bool {
		if matchingBooks[i].Authors != matchingBooks[j].Authors {
			return matchingBooks[i].Authors < matchingBooks[j].Authors
		}
		return matchingBooks[i].Title < matchingBooks[j].Title
	})

	return matchingBooks
}
