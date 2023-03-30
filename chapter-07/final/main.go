package main

import (
	"fmt"
	"os"

	"learngo-pockets/genericworms/books"
	"learngo-pockets/genericworms/collectors"
)

func main() {
	bookworms, err := collectors.Load[books.Book]("books/testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := bookworms.FindCommon()

	fmt.Println("Here are the common books:")
	displayBooks(commonBooks)
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []books.Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
