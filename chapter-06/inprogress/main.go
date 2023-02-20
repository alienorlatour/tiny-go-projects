package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	matchingBooks := findMatchingBooks(bookworms)

	fmt.Println("Here are the matching books:")
	displayBooks(matchingBooks)

	suggestions := suggestOtherBooks(bookworms)
	displaySuggestions(suggestions)
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}

func displaySuggestions(suggestionsForBookworms []Bookworm) {
	for _, bookworm := range suggestionsForBookworms {
		fmt.Printf("\nHere are the suggestions for %s:\n", bookworm.Name)
		displayBooks(bookworm.Books)
		fmt.Println()
	}
}
