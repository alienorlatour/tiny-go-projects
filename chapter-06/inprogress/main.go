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

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are the common books:")
	displayBooks(commonBooks)

	recommendations := recommendOtherBooks(bookworms)
	displayRecommendations(recommendations)
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}

func displayRecommendations(recommendations []Bookworm) {
	for _, bookworm := range recommendations {
		fmt.Printf("\nHere are the recommendations for %s:\n", bookworm.Name)
		displayBooks(bookworm.Books)
		fmt.Println()
	}
}
