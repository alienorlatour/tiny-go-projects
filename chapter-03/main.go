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

	sol := newSolution([]byte("slice"))
	reader := bufio.NewReader(os.Stdin)
	for {
		attempt := input(reader)
		if bytes.Equal(attempt, sol.word) {
			// win
			fmt.Println("Bravo! You found the word.")
			return
		}

		sol.feedback(attempt)
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
type solution struct {
	word      []byte
	positions map[byte][]int
}

func newSolution(word []byte) solution {
	sol := solution{
		word:      word,
		positions: make(map[byte][]int),
	}

	return sol
}

// prints out hints on how to find the solution
func (s *solution) feedback(attempt []byte) []status {
	for i, letter := range s.word {
		// appending to a nil slice will return a slice, this is safe
		s.positions[letter] = append(s.positions[letter], i)
	}

	f := make([]status, wordLength)
	// scan the attempts and check if they are in the solution
	for i, letter := range attempt {
		// keep track of already seen characters
		correctness := s.checkLetter(letter, i)
		if correctness == correctPosition {
			// remove found letter from positions
			s.markLetterAsSeen(letter, i)
			f[i] = correctPosition
		}
	}

	for i, letter := range attempt {
		correctness := s.checkLetter(letter, i)
		if correctness == wrongPosition {
			// remove the left-most occurrence
			s.positions[letter] = s.positions[letter][1:]
		}
		f[i] = correctness
	}

	return f
}

func (s *solution) markLetterAsSeen(letter byte, positionInWord int) {
	positions := s.positions[letter]

	for i, pos := range positions {
		if pos == positionInWord {
			s.positions[letter] = append(positions[:i], positions[i:]...)
		}
	}
}

// checkLetter returns the correctness of a letter
// at the specified index in the solution.
func (s *solution) checkLetter(letter byte, index int) status {
	positions, ok := s.positions[letter]
	if !ok || len(positions) == 0 {
		return absentCharacter
	}

	for _, pos := range positions {
		if pos == index {
			return correctPosition
		}
	}

	return wrongPosition
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
