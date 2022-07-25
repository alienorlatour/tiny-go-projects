package gordle

import (
	"bufio"
	"fmt"
	"os"
)

const wordLength = 5

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	reader lineReader
}

// New returns a Gordle variable, which can be used to Play!
func New() *Gordle {
	g := &Gordle{
		reader: bufio.NewReader(os.Stdin),
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
		attempt        []byte
		attemptIsValid bool
		err            error
	)

	for !attemptIsValid {
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
		} else {
			attemptIsValid = true
		}
	}

	return string(attempt)
}

var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of letters as the solution ")

func (g Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != wordLength {
		return fmt.Errorf("expected %d, got %d, %w", wordLength, len(attempt), errInvalidWordLength)
	}

	return nil
}
