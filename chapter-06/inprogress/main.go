package main

import "fmt"

func main() {
	readers, err := loadLectors("testdata/lectors.json")
	if err != nil {
		panic(err)
	}

	matchingBooks := findMatchingBooks(readers)

	dsiplayMatchingBooks(matchingBooks)
}

func dsiplayMatchingBooks(books []Book) {
	fmt.Println("Here are the matching books:")
	for _, book := range books {
		fmt.Println(book.Title, "by", book.Authors)
	}
}
