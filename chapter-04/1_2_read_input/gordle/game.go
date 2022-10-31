package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const wordLength = 5

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader *bufio.Reader
}

// New returns a Game variable, which can be used to Play!
func New(reader io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(reader),
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	// ask for a valid word
	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
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
		if len(attempt) != wordLength {
			_, _ = fmt.Fprintf(os.Stderr, "invalid word length: expected %d, got %d\n", wordLength, len(attempt))
		} else {
			return attempt
		}
	}
}
