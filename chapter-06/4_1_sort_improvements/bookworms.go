package main

import (
	"encoding/json"
	"os"
	"sort"
)

// A Bookworm contains the list of books on a bookworm's shelf.
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
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var bookworms []Bookworm

	// Decode the file and store the content in the variable bookworms.
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// findCommonBooks returns books that are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	// Register all books on shelves.
	booksOnShelves := booksCount(bookworms)

	// List containing all the books that were read by at least 2 bookworms.
	var commonBooks []Book

	// Find books that were added to shelve more than once.
	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	// Sort allows us to be deterministic, sorted alphabetically by authors and then by title.
	return sortBooks(commonBooks)
}

// booksCount registers all the books and their occurrences from the bookworms shelves.
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			// If a bookworm has two copies, that counts as two.
			count[book]++
		}
	}

	return count
}

// byAuthor is a list of Book. Defining a custom type to implement sort.Interface
type byAuthor []Book

// Len implements sort.Interface by returning the length of BookByAuthor.
func (b byAuthor) Len() int { return len(b) }

// Swap implements sort.Interface and swaps two books.
func (b byAuthor) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements sort.Interface and returns BookByAuthor sorted by Author and then Title.
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

// sortBooks sorts the books by Author and then Title in alphabetical order.
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}
