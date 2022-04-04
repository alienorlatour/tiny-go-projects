package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// all words in the corpus have this many letters
	wordLength = 5
)

var (
	errInvalidWordLength = fmt.Errorf("word should be %d letters", wordLength)
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
			_, _ = fmt.Fprintf(os.Stderr, "error while reading the player's input: %s", err.Error())
			continue
		}

		err = validate(attempt)
		if err != nil {
			fmt.Println(err)
		} else {
			attemptIsValid = true
		}
	}

	return string(attempt)
}

func validate(attempt []byte) error {
	if len(attempt) != wordLength {
		return errInvalidWordLength
	}

	return nil
}
