package gordle

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	reader          *bufio.Reader
	solution        []rune
	maxAttempts     int
	solutionChecker *solutionChecker
}

// New returns a Gordle variable, which can be used to Play!
func New(corpus []string, cfs ...ConfigFunc) (*Gordle, error) {
	g := &Gordle{
		reader:      bufio.NewReader(os.Stdin), // read from stdin by default
		maxAttempts: -1,                        // no maximum number of attempts by default
		solution:    pickWord(corpus),          // pick a random word from the corpus
	}

	// Apply the configuration functions after defining the default values, as they override them.
	for _, cf := range cfs {
		err := cf(g)
		if err != nil {
			return nil, fmt.Errorf("unable to apply config func: %w", err)
		}
	}

	// Delay the checker creation till here, in case the solution was passed as a config func.
	g.solutionChecker = &solutionChecker{solution: g.solution}
	return g, nil
}

// Play runs the game.
func (g *Gordle) Play() {
	// break condition: we've reached the maximum number of attempts
	for currentAttempt := 0; currentAttempt < g.maxAttempts; currentAttempt++ {
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
func (g *Gordle) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		// Read the attempt from the player.
		suggestion, _, err := g.reader.ReadLine()
		if err != nil {
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

// errInvalidWordLength is only used internally
var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of characters as the solution ")

// validateAttempt ensures the attempt is valid enough.
func (g *Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(attempt), errInvalidWordLength)
	}

	return nil
}
