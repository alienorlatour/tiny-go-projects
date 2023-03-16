package books

import (
	"encoding/json"
	"os"
	"sort"

	"learngo-pockets/genericworms/collectors"
)

// Book describes an item on a collector's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// Collectors describe a list of book collectors and their books
type Collectors collectors.Collectors[Book]

// Load reads the file and returns the list of collectors, and their beloved books, found therein.
func Load(filePath string) (Collectors, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var colls Collectors

	// Decode the file and store the content in the variable colls.
	err = json.NewDecoder(f).Decode(&colls)
	if err != nil {
		return nil, err
	}

	return colls, nil
}

// FindCommon return the books in common, sorted first by author and then title.
func (colls Collectors) FindCommon() []Book {
	// We need a Collectors[T] here
	booksInCommon := collectors.Collectors[Book](colls).FindCommon()

	// sort.Sort sorts the slice in place.
	// We can instantiate a slice of the type byAuthor, which implements sort.Interface.
	sort.Sort(byAuthor(booksInCommon))

	return booksInCommon
}

// byAuthor implements sort.Interface for a list of books.
type byAuthor []Book

// Len implements sort.Interface by returning the length of Books.
func (b byAuthor) Len() int { return len(b) }

// Less returns true if Author i is before Author j in alphabetical order.
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

// Swap implements sort.Interface and swaps two books.
func (b byAuthor) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
