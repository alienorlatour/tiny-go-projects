package main

import (
	"fmt"
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

// Books is a list of Book. Defining a custom type to implement sort.Interface
type Books []Book

// String implements Stringer interface.
func (b Books) String() {
	for _, book := range b {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}

// Len implements sort.Interface by returning the length of Books.
func (b Books) Len() int { return len(b) }

// Swap implements sort.Interface and swaps two books.
func (b Books) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements sort.Interface and returns Books sorted by Author and then Title.
func (b Books) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b.LessByAuthor(i, j)
	}
	return b.LessByTitle(i, j)
}

// LessByAuthor returns true if Author i is before Author j in alphabetical order.
func (b Books) LessByAuthor(i, j int) bool {
	return b[i].Author < b[j].Author

}

// LessByTitle returns true if Title i is before Title j in alphabetical order.
func (b Books) LessByTitle(i, j int) bool {
	return b[i].Title < b[j].Title
}

// sortBooks sorts the books by Author and then Title in alphabetical order.
func sortBooks(books []Book) []Book {
	sort.Sort(Books(books))
	return books
}
