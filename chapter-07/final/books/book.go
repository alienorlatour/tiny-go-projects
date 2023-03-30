package books

import (
	"fmt"
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

// FindCommon return the books in common, sorted first by author and then title.
func (colls Collectors) FindCommon() []Book {
	// We need a Collectors[T] here
	booksInCommon := collectors.Collectors[Book](colls).FindCommon()

	// sort.Sort sorts the slice in place.
	// We can instantiate a slice of the type byAuthor, which implements sort.Interface.
	sort.Sort(byAuthor(booksInCommon))

	return booksInCommon
}

// Display prints out the titles and authors of a list of books
func Display(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
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
