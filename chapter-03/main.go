package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-03/wordle"
)

const (
	// all words in the corpus have this many letters
	wordLength = 5
)

func main() {
	fmt.Println("Welcome to Gordle!")

	sol := wordle.NewSolution([]byte("slice"))
	reader := bufio.NewReader(os.Stdin)
	for {
		attempt := askWord(reader)
		if sol.IsWord(attempt) {
			// win
			fmt.Println("Bravo! You found the word.")
			return
		}

		f := sol.Feedback(attempt)
		fmt.Println(f)
	}

}

type lineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// askWord prints out the instruction and reads from the standard askWord
func askWord(reader lineReader) []byte {
	fmt.Println("Enter a guess:")

	var attempt []byte
	var attemptIsValid bool
	var err error

	for !attemptIsValid {
		attempt, _, err = reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while reading the player's askWord: %s", err.Error())
			continue
		}

		err = validateInput(attempt)
		if err != nil {
			fmt.Println(err)
		} else {
			attemptIsValid = true
		}
	}

	return attempt
}

var (
	errInvalidWordLength = fmt.Errorf("word has the wrong number of characters")
)

func validateInput(attempt []byte) error {
	if len(attempt) != wordLength {
		return errInvalidWordLength
	}

	return nil
}
