package gordle

import (
	"fmt"
	"os"
)

const wordLength = 5

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
}

// New returns a Gordle variable, which can be used to Play!
func New() *Gordle {
	g := &Gordle{}

	return g
}

// Play runs the game.
func (g *Gordle) Play() {
	fmt.Printf("Enter a %d-letter guess:\n", wordLength)

	var (
		attempt        []rune
		attemptIsValid bool
	)

	for !attemptIsValid {
		// Read the attempt from the player.
		wordCount, err := fmt.Fscanf(os.Stdin, "%s", &attempt)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while reading the player's word: %q", err)
			continue
		}
		// We expect a single word.
		if wordCount != 1 {
			fmt.Fprintf(os.Stderr, "error while reading the player's word: a single word wasn't provided, got %d instead", wordCount)
			continue
		}

		// Verify the suggestion has a valid length.
		err = g.validateAttempt(attempt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		} else {
			attemptIsValid = true
		}
	}

	fmt.Printf("Your guess: %q\n", attempt)
}

var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of letters as the solution ")

func (g Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != wordLength {
		return fmt.Errorf("expected %d, got %d, %w", wordLength, len(attempt), errInvalidWordLength)
	}

	return nil
}
