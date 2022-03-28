package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	wordLength = 5
)

func main() {
	fmt.Println("Welcome to Gordle!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(input(reader))
}

type lineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// input prints out the instruction and reads from the standard input
func input(reader lineReader) string {
	fmt.Println("Enter a guess:")

	var attempt []byte
	var attemptIsValid bool
	var err error

	for !attemptIsValid {
		attempt, _, err = reader.ReadLine()
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}

		if len(attempt) != wordLength {
			fmt.Printf("word %q should be %d letters, got: %d\n", attempt, wordLength, len(attempt))
		} else {
			attemptIsValid = true
		}
	}

	return string(attempt)
}
