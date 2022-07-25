package gordle

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const wordLength = 5

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	reader   lineReader
	solution []rune
}

// New returns a Gordle variable, which can be used to Play!
func New(solution []rune) *Gordle {
	g := &Gordle{
		reader:   bufio.NewReader(os.Stdin),
		solution: solution,
	}

	return g
}

type lineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// Play runs the game.
func (g *Gordle) Play() string {
	fmt.Printf("Enter a %d-letter guess:\n", wordLength)

	var (
		attempt []byte
		err     error
	)

	for {
		// Read the attempt from the player.
		attempt, _, err = g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		attemptRunes := []rune(string(attempt))
		// Verify the suggestion has a valid length.
		err = g.validateAttempt(attemptRunes)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}

		solutionBytes := []byte(string(g.solution))
		if bytes.Equal(attempt, solutionBytes) {
			// win
			fmt.Println("Bravo! You found the word.")
			return string(attempt)
		} else {
			fmt.Println("Give it again a try:")
		}
	}
}

var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of letters as the solution ")

func (g Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != wordLength {
		return fmt.Errorf("expected %d, got %d, %w", wordLength, len(attempt), errInvalidWordLength)
	}

	return nil
}
