package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"tiny-go-projects/chapter-03/wordle"
)

//go:embed corpus.txt
var corpus string

const (
	// all words in the corpus have this many letters
	wordLength = 5
)

func main() {
	fmt.Println("Welcome to Gordle!")

	sol := wordle.NewSolution(pickOne(corpus))
	reader := bufio.NewReader(os.Stdin)
	nbTries := 0

	for {
		attempt := askWord(reader)
		nbTries++
		if sol.IsWord(attempt) {
			// win
			fmt.Printf("Bravo! You found the word in %d attempts.\n", nbTries)
			return
		}

		f := sol.Feedback(attempt)
		fmt.Println(f)
	}

}

// pickOne returns a random word from the corpus
func pickOne(corpus string) []byte {
	list := strings.Split(corpus, "\n")
	index := rand.Int() % len(list)

	word := strings.ToUpper(list[index])
	return []byte(word)
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

		attempt = []byte(strings.ToUpper(string(attempt)))
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
	errInvalidWordLength = fmt.Errorf("word has the wrong number of characters, try again")
)

func validateInput(attempt []byte) error {
	if len(attempt) != wordLength {
		return errInvalidWordLength
	}

	return nil
}
