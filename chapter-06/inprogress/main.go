package main

import "fmt"

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		panic(err)
	}

	matchingBooks := findMatchingBooks(bookworms)

	fmt.Println("Here are the matching books:")
	displayBooks(matchingBooks)

	fmt.Println("")

	suggestions := suggestOtherBooks(bookworms)
	displaySuggestions(suggestions)
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println(book.Title, "by", book.Authors)
	}
}

func displaySuggestions(suggestionsForBookworms []Bookworm) {
	for _, bookworm := range suggestionsForBookworms {
		fmt.Printf("Here are the suggestions for %s:\n", bookworm.Name)
		displayBooks(bookworm.Books)
		fmt.Println("")
	}
}
