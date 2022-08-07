package gordle

import (
	"bufio"
	"fmt"
	"os"
)

const wordLength = 5

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	reader          *bufio.Reader
	solution        []rune
	maxAttempts     int
	currentAttempt  int
	solutionChecker *solutionChecker
}

// New returns a Gordle variable, which can be used to Play!
func New(reader *bufio.Reader, solution []rune, maxAttempts int) *Gordle {
	g := &Gordle{
		reader:      reader,
		solution:    solution,
		maxAttempts: maxAttempts,
	}

	g.solutionChecker = &solutionChecker{solution: g.solution}

	return g
}

// Play runs the game.
func (g *Gordle) Play() {
	// break condition: we've reached the maximum number of attempts
	for g.currentAttempt != g.maxAttempts {
		// ask for a valid word
		attempt := g.ask()

		// check it
		fb := g.solutionChecker.check(attempt)

		// print the feedback
		fmt.Println(fb.String())

		if string(attempt) == string(g.solution) {
			fmt.Printf("🎉 You won! You found in %d attempt(s)! The word was: %s.\n", g.currentAttempt, string(g.solution))
			return
		}
		g.currentAttempt++
	}

	// we've exhausted the number of allowed attempts
	fmt.Printf("😞 You've lost! The solution was: %s. \n", string(g.solution))
}

// ask reads input until a valid suggestion is made (and returned).
func (g Gordle) ask() []rune {
	fmt.Printf("Enter a %d-letter guess:\n", wordLength)

	for {
		// Read the attempt from the player.
		suggestion, _, err := g.reader.ReadLine()
		if err != nil {
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

var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of letters as the solution ")

// validateAttempt ensures the attempt is valid enough.
func (g Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != wordLength {
		return fmt.Errorf("expected %d, got %d, %w", wordLength, len(attempt), errInvalidWordLength)
	}

	return nil
}
