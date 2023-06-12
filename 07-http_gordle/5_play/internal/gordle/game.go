package gordle

import (
	"fmt"
	"os"
	"strings"
)

// Game holds the information we need to get the Feedback of a play.
type Game struct {
	solution []rune
}

// New returns a Game variable, which can be used to Play!
func New(solution string) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrEmptyCorpus
	}

	return &Game{
		solution: splitToUppercaseCharacters(solution),
	}, nil
}

const (
	// ErrInvalidGuessLength indicates a guess doesn't have the right number of characters.
	ErrInvalidGuessLength = gameError("invalid guess length")
)

// Play runs the game. If the guess is not valid, we return ErrInvalidGuessLength.
func (g *Game) Play(guess string) (Feedback, error) {
	err := g.validateGuess(guess)
	if err != nil {
		return Feedback{}, fmt.Errorf("this guess is not the correct length: %w", err)
	}

	// check it
	characters := splitToUppercaseCharacters(guess)
	fb := computeFeedback(characters, g.solution)
	return fb, nil
}

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess string) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("you guessed a %d word length, remember the answer is %d word length, %w", len(guess), len(g.solution), ErrInvalidGuessLength)
	}

	return nil
}

// ShowAnswer gives up on playing this game. It returns the solution.
func (g *Game) ShowAnswer() string {
	return string(g.solution)
}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters.
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// computeFeedback verifies every character of the guess against the solution.
func computeFeedback(guess, solution []rune) Feedback {
	// initialise holders for marks
	result := make(Feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "guess and solution have different lengths: %d vs %d", len(guess), len(solution))
		// return a Feedback full of absent characters
		return result
	}

	// check for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	// look for letters in the wrong position
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			// The character has already been marked, ignore it.
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// The letter of the solution is already assigned to a letter of the guess.
				// Skip to the next letter of the solution.
				continue
			}

			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				// Skip to the next letter of the guess.
				break
			}
		}
	}

	return result
}
