package main

import (
	"bufio"
	"bytes"
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

	sol := []byte("slice")

	reader := bufio.NewReader(os.Stdin)
	for {
		attempt := input(reader)
		if bytes.Equal(attempt, sol) {
			// win
			fmt.Println("Bravo! You found the word.")
			return
		}

		feedback(attempt, newSolution(sol))
	}

}

type status int

const (
	correctPosition status = iota
	wrongPosition
	absentCharacter
)

// solution holds the positions of the valid characters
// since a single character can appear several times, we store these times as a slice of indexes
type solution map[byte][]int

func newSolution(word []byte) solution {
	sol := solution{}
	for i, letter := range word {
		//appending to a nil slice will return a slice, this is safe
		sol[letter] = append(sol[letter], i)
	}
	return sol
}

// prints out hints on how to find the solution
func feedback(attempt []byte, s solution) []status {
	f := make([]status, wordLength)
	// scan the attempts and check if they are in the solution
	for i, letter := range attempt {
		// keep track of already seen characters
		f[i] = s.checkLetter(letter, i)
	}
	return f
}

func (s solution) checkLetter(letter byte, index int) status {
	positions, ok := s[letter]
	if !ok {
		return absentCharacter
	}

	for _, pos := range positions {
		if pos == index {
			return correctPosition
		}
	}

	return absentCharacter
}

type lineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// input prints out the instruction and reads from the standard input
func input(reader lineReader) []byte {
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

	return attempt
}

func validate(attempt []byte) error {
	if len(attempt) != wordLength {
		return errInvalidWordLength
	}

	return nil
}
