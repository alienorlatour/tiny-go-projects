package gordle

import (
	"bufio"
	"fmt"
	"os"
)

const wordLength = 5

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	reader *bufio.Reader
}

// New returns a Gordle variable, which can be used to Play!
func New(reader *bufio.Reader) *Gordle {
	g := &Gordle{
		reader: reader,
	}
	return g
}

// Play runs the game.
func (g *Gordle) Play() {

	// ask for a valid word
	attempt := g.ask()

	fmt.Printf("Your guess is: %s\n", string(attempt))
}

// ask reads input until a valid suggestion is made (and returned).
func (g Gordle) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", wordLength)

	for {
		// Read the attempt from the player.
		suggestion, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		attempt := []rune(string(suggestion))

		// Verify the suggestion has a valid length.
		if len(attempt) != wordLength {
			_, _ = fmt.Fprintf(os.Stderr, "invalid word length: expected %d, got %d\n", wordLength, len(attempt))
			continue
		} else {
			return attempt
		}
	}
}
