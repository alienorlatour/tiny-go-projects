package gordle

import (
	"fmt"
	"io"
	"os"
)

const (
	wordLength = 5
)

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	reader io.Reader
}

// New returns a Gordle variable, which can be used to Play!
func New() *Gordle {
	g := &Gordle{
		reader: os.Stdin, // read from stdin by default
	}

	return g
}

// Play runs the game.
func (g *Gordle) Play() {
	fmt.Printf("Enter a %d-letter guess:\n", wordLength)

	var (
		suggestion     []byte
		attemptIsValid bool
	)

	for !attemptIsValid {
		// Read the suggestion from the player.
		wordCount, err := fmt.Fscanf(g.reader, "%s", &suggestion)
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
		if len(suggestion) != wordLength {
			fmt.Printf("word %q should be %d letters, got: %d\n", suggestion, wordLength, len(suggestion))
		} else {
			attemptIsValid = true
		}
	}

	fmt.Printf("Your guess: %s\n", suggestion)
}
