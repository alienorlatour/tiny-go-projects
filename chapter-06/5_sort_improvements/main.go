package main

import "fmt"

func main() {
	handmaidsTale := Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake := Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar := Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	ilPrincipe := Book{Author: "Niccolò Machiavelli", Title: "Il Principe"}

	// Create a list of books.
	books := Books{theBellJar, handmaidsTale, ilPrincipe, oryxAndCrake}
	// Print the initial list of books.
	fmt.Println("Unsorted books:")
	books.String()

	fmt.Println("---")
	fmt.Println("Sorted books:")
	// Sort the books by Author and then Title.
	sortBooks(books)
	// Print the sorted list of books.
	books.String()
}
