package main

import (
	"fmt"
	"os"

	"learngo-pockets/genericworms/books"
	"learngo-pockets/genericworms/collectors"
)

func main() {
	bookworms, err := collectors.Load[books.Book]("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := bookworms.FindCommon()

	fmt.Println("Here are the common books:")
	books.Display(commonBooks)

	// TODO: Test patterns
}
