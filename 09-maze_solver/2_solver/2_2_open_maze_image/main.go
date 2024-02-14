package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)

	_, err := openMaze(inputFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}

// usage displays the usage of the binary and exits the program.
func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}
