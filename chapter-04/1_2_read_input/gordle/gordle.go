package gordle

import (
	"bufio"
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
		attempt        []byte
		attemptIsValid bool
		err            error
		reader         = bufio.NewReader(os.Stdin)
	)

	for !attemptIsValid {
		// Read the attempt from the player.
		attempt, _, err = reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		// Verify the suggestion has a valid length.
		if len(attempt) != wordLength {
			_, _ = fmt.Fprintf(os.Stderr, "invalid word length: expected %d, got %d\n", wordLength, len(attempt))
		} else {
			attemptIsValid = true
		}
	}

	fmt.Printf("Your guess: %q\n", attempt)
}
