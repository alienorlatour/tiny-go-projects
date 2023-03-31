package main

import (
	"fmt"
	"os"

	"learngo-pockets/genericworms/books"
	"learngo-pockets/genericworms/collectors"
	"learngo-pockets/genericworms/patterns"
)

func main() {
	bookworms, err := collectors.Load[books.Book]("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := bookworms.FindCommon()

	fmt.Println("Here are the common books:")
	books.Display(os.Stdout, commonBooks)

	crafters, err := collectors.Load[patterns.Pattern]("testdata/patterns.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load crafters: %s\n", err)
		os.Exit(1)
	}

	commonPatterns := crafters.FindCommon()

	fmt.Println("Here are the common patterns:")
	patterns.Display(os.Stdout, commonPatterns)
}
