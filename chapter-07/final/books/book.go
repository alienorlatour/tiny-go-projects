package books

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"learngo-pockets/genericworms/collectors"
)

// Book describes an item on a collectors's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func (book Book) String() string {
	return fmt.Sprintf("%s: %s", book.Author, book.Title)
}

// Load reads the file and returns the list of collectors, and their beloved books, found therein.
func Load(filePath string) (collectors.Collectors[Book], error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var colls []collectors.Collector[Book]

	// Decode the file and store the content in the variable colls.
	err = json.NewDecoder(f).Decode(&colls)
	if err != nil {
		return nil, err
	}

	return colls, nil
}

// sortBooks sorts the books by Author and then Title.
func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})

	return books
}

func (book Book) Less(other Book) bool {
	if book.Author != other.Author {
		return book.Author < other.Author
	}
	return book.Title < other.Title
}
