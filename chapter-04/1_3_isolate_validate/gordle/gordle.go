package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const wordLength = 5

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	reader *bufio.Reader
}

// New returns a Gordle variable, which can be used to Play!
func New(reader io.Reader) *Gordle {
	g := &Gordle{
		reader: bufio.NewReader(reader),
	}
	fmt.Println("Welcome to Gordle!")

	return g
}

// Play runs the game.
func (g *Gordle) Play() {

	// ask for a valid word
	attempt := g.ask()

	fmt.Printf("Your guess is: %s\n", string(attempt))
}

// ask reads input until a valid suggestion is made (and returned).
func (g *Gordle) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", wordLength)

	for {
		// Read the attempt from the player.
		suggestion, _, err := g.reader.ReadLine()
		if err != nil {
			// We failed to read this line, maybe the next one is better?
			// Let’s give it a chance.
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		attempt := []rune(string(suggestion))

		// Verify the suggestion has a valid length.
		err = g.validateAttempt(attempt)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		} else {
			return attempt
		}
	}
}

// errInvalidWordLength
var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of characters as the solution ")

// validateAttempt ensures the attempt is valid enough.
func (g *Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != wordLength {
		return fmt.Errorf("expected %d, got %d, %w", wordLength, len(attempt), errInvalidWordLength)
	}

	return nil
}
