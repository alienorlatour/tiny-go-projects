package main

import "fmt"

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		panic(err)
	}

	matchingBooks := findMatchingBooks(bookworms)

	displayMatchingBooks(matchingBooks)
}

func displayMatchingBooks(books []Book) {
	fmt.Println("Here are the matching books:")
	for _, book := range books {
		fmt.Println(book.Title, "by", book.Authors)
	}
}
