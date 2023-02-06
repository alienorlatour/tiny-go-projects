package main

import (
	"encoding/json"
	"os"
	"sort"
)

// A Bookworm contains the list of books of a specific person.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms, and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {
	// Open the file to get an io.Reader.
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Initialise the type in which the file will be decoded.
	bookworms := make([]Bookworm, 0)

	// Decode the file and stores the content in the value bookworms.
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// findMatchingBooks returns books that are on more than one bookworm's shelf.
func findMatchingBooks(bookworms []Bookworm) []Book {
	booksOnShelves := make(map[Book]uint)

	// Register all books on shelves.
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			booksOnShelves[book]++
		}
	}

	// List containing all the books that were read by at least 2 bookworms.
	booksReadBySeveralBookworms := make([]Book, 0)

	// Find books that were added to shelves more than once.
	for book, count := range booksOnShelves {
		if count > 1 {
			booksReadBySeveralBookworms = append(booksReadBySeveralBookworms, book)
		}
	}

	// sort allows us to be deterministic, sorted alphabetically by authors and then by title.
	sort.Slice(booksReadBySeveralBookworms, func(i, j int) bool {
		if booksReadBySeveralBookworms[i].Author != booksReadBySeveralBookworms[j].Author {
			return booksReadBySeveralBookworms[i].Author < booksReadBySeveralBookworms[j].Author
		}
		return booksReadBySeveralBookworms[i].Title < booksReadBySeveralBookworms[j].Title
	})

	return booksReadBySeveralBookworms
}
