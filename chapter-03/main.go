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
	fmt.Println(input())
}

// input prints out the instruction and reads from the standard input
func input() string {
	fmt.Println("Enter a guess:")

	var attempt []byte
	var attemptIsValid bool
	var err error
	var reader = bufio.NewReader(os.Stdin)

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
