package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

// Game holds all the information we need to play a game of Gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a Game variable, which can be used to Play!
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	// break condition: we've reached the maximum number of attempts
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		// ask for a valid word
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}

	// we've exhausted the number of allowed attempts
	fmt.Printf("😞 You've lost! The solution was: %s. \n", string(g.solution))
}

// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		// Read the guess from the player.
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			// We failed to read this line, maybe the next one is better?
			// Let’s give it a chance.
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		// Verify the suggestion has a valid length.
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters.
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
