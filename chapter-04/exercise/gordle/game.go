package gordle

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/text/unicode/norm"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader          *bufio.Reader
	solution        []rune
	maxAttempts     int
	solutionChecker *solutionChecker
}

// New returns a Game variable, which can be used to Play!
func New(corpus []string, cfs ...ConfigFunc) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	g := &Game{
		reader:      bufio.NewReader(os.Stdin),                    // read from stdin by default
		solution:    splitToUppercaseCharacters(pickWord(corpus)), // pick a random word from the corpus
		maxAttempts: -1,                                           // no maximum number of attempts by default
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
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	// break condition: we've reached the maximum number of attempts
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		// ask for a valid word
		guess := g.ask()

		// check it
		fb := g.solutionChecker.check(guess)

		// print the feedback
		fmt.Println(fb.String())

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
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		// Verify the suggestion has a valid length.
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
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
	iter := norm.Iter{}
	iter.Init(norm.NFKC, []byte(input))

	return []rune(strings.ToUpper(input))
}
