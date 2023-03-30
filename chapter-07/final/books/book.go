package books

import (
	"fmt"
	"io"
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

	// sort.Slice sorts the slice in place.
	sort.Slice(booksInCommon, func(i, j int) bool {
		if booksInCommon[i].Author != booksInCommon[j].Author {
			return booksInCommon[i].Author < booksInCommon[j].Author
		}
		return booksInCommon[i].Title < booksInCommon[j].Title
	})

	return booksInCommon
}

// Display prints out the titles and authors of a list of books
func Display(w io.Writer, books []Book) {
	for _, book := range books {
		_, _ = fmt.Fprintln(w, "-", book.Title, "by", book.Author)
	}
}
