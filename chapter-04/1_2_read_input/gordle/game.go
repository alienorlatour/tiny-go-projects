package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const solutionLength = 5

// Game holds all the information we need to play a game of Gordle.
type Game struct {
	reader *bufio.Reader
}

// New returns a Game variable, which can be used to Play!
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
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
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		// Read the guess from the player.
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			// We failed to read this line, maybe the next one is better?
			// Let’s give it a chance.
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := []rune(string(playerInput))

		// Verify the suggestion has a valid length.
		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution! Expected %d characters, got %d.\n", solutionLength, len(guess))
		} else {
			return guess
		}
	}
}
