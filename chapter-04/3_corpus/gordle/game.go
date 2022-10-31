package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader          *bufio.Reader
	solution        []rune
	maxAttempts     int
	solutionChecker *solutionChecker
}

// New returns a Game variable, which can be used to Play!
func New(reader io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}
	g := &Game{
		reader:      bufio.NewReader(reader),
		solution:    []rune(strings.ToUpper(pickWord(corpus))), // pick a random word from the corpus
		maxAttempts: maxAttempts,
	}

	g.solutionChecker = &solutionChecker{solution: g.solution}

	return g, nil
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordlez<!")

	// break condition: we've reached the maximum number of attempts
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		// ask for a valid word
		attempt := g.ask()

		// check it
		fb := g.solutionChecker.check(attempt)

		// print the feedback
		fmt.Println(fb.String())

		if string(attempt) == string(g.solution) {
			fmt.Printf("🎉 You won! You found in %d attempt(s)! The word was: %s.\n", currentAttempt, string(g.solution))
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
		// Read the attempt from the player.
		suggestion, _, err := g.reader.ReadLine()
		if err != nil {
			// We failed to read this line, maybe the next one is better?
			// Let’s give it a chance.
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		// Verify the suggestion has a valid length.
		attempt := []rune(strings.ToUpper(string(suggestion)))
		err = g.validateAttempt(attempt)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		} else {
			return attempt
		}
	}
}

var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of characters as the solution ")

// validateAttempt ensures the attempt is valid enough.
func (g *Game) validateAttempt(attempt []rune) error {
	if len(attempt) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(attempt), errInvalidWordLength)
	}

	return nil
}
