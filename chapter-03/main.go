package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"tiny-go-projects/chapter-03/wordle"
)

//go:embed corpus_5letters.txt
var corpus string

const (
	// all words in the corpus have this many letters
	wordLength = 5
	// the number of attempts the player has to find the word
	maxTries = 6
)

func main() {
	fmt.Println("Welcome to Gordle!")

	solution := pickOne(corpus)
	sol := wordle.NewSolution(solution)
	reader := runeReader{
		byteReader: bufio.NewReader(os.Stdin),
	}

	nbTries := 0

	for {
		attempt := askWord(reader)
		nbTries++
		if nbTries == maxTries {
			fmt.Printf("😞 You lost! The solution was: %s. \n", string(solution))
			return
		}

		if sol.IsWord(attempt) {
			// win
			fmt.Printf("Bravo! You found the word in %d attempts.\n", nbTries)
			return
		}

		f := sol.Feedback(attempt)
		fmt.Println(wordle.StatusesToString(f))
	}
}

// pickOne returns a random word from the corpus
func pickOne(corpus string) []rune {
	list := strings.Split(corpus, "\n")

	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Int() % len(list)

	word := strings.ToUpper(list[index])
	return []rune(word)
}

type lineReader interface {
	ReadLine() ([]rune, error)
}

type runeReader struct {
	byteReader *bufio.Reader
}

// ReadLine reads a line of runes
func (r runeReader) ReadLine() ([]rune, error) {
	bytes, _, err := r.byteReader.ReadLine()
	if err != nil {
		return nil, err
	}

	return []rune(string(bytes)), nil
}

// askWord prints out the instruction and reads from the standard askWord
func askWord(reader lineReader) []rune {
	fmt.Println("Enter a guess:")

	var attempt []rune
	var attemptIsValid bool
	var err error

	for !attemptIsValid {
		attempt, err = reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while reading the player's askWord: %s", err.Error())
			continue
		}

		attempt = []rune(strings.ToUpper(string(attempt)))
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
	errInvalidWordLength = fmt.Errorf("Word has the wrong number of characters, try again:")
)

func validateInput(attempt []rune) error {
	if len(attempt) != wordLength {
		return errInvalidWordLength
	}

	return nil
}
