package gordle

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// New returns a Gordle variable, which can be used to Play!
func New(cfs ...ConfigFunc) (*Gordle, error) {
	g := &Gordle{
		scanner:     bufio.NewScanner(os.Stdin), // read from stdin by default
		maxAttempts: -1,                         // no maximum number of attempts by default
		solution:    randomWord(),               // pick a random word from the corpus
	}

	for _, cf := range cfs {
		err := cf(g)
		if err != nil {
			return nil, fmt.Errorf("unable to apply config func: %w", err)
		}
	}

	return g, nil
}

// Play runs the game. It will exit when the maximum number of attempts was reached, or if the word was found.
func (g *Gordle) Play() {
	// break condition: we've reached the maximum number of attempts
	for g.currentAttempt != g.maxAttempts {
		// ask for a valid word
		word := g.ask()

		// check it
		fb := g.checkAgainstSolution(word)

		// print the feedback
		fmt.Println(fb)
		if string(word) == string(g.solution) {
			fmt.Printf("🎉 You won! You found in %d attempt(s)! The word was: %s.\n", g.currentAttempt, string(g.solution))
			return
		}
		g.currentAttempt++
	}
	// we've exhausted the number of allowed attempts
	fmt.Printf("😞 You've lost! The solution was: %s. \n", string(g.solution))
}

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	scanner        scanner
	solution       []rune
	maxAttempts    int
	currentAttempt int
	positions      map[rune][]int
}

// ask scans until a valid suggestion is made (and returned).
func (g *Gordle) ask() []rune {
	fmt.Println("Enter a guess:")

	for g.scanner.Scan() {
		suggestion := g.scanner.Text()
		if g.scanner.Err() != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while reading the player's word: %q", g.scanner.Err())
			continue
		}

		attempt := []rune(strings.ToUpper(suggestion))
		err := g.validateAttempt(attempt)
		if err != nil {
			fmt.Println(err)
		} else {
			return attempt
		}
	}
	// this can't happen
	return []rune{}
}

var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of letters as the solution ")

// validateAttempt ensures the attempt is valid enough.
func (g *Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != len(g.solution) {
		return errInvalidWordLength
	}
	return nil
}

// checkAgainstSolution checks every letter of the word against the solution.
func (g *Gordle) checkAgainstSolution(word []rune) feedback {
	// reset the positions map
	g.positions = make(map[rune][]int)

	for i, letter := range g.solution {
		// appending to a nil slice will return a slice, this is safe
		g.positions[letter] = append(g.positions[letter], i)
	}

	fb := make(feedback, len(g.solution))

	// scan the attempts and check if they are in the solution
	for i, letter := range word {
		// keep track of already seen characters
		correctness := g.checkLetterAtPosition(letter, i)
		if correctness == correctPosition {
			// remove found letter from positions
			g.markLetterAsSeen(letter, i)
			fb[i] = correctPosition
		}
	}

	for i, letter := range word {
		if fb[i] == correctPosition {
			continue
		}

		correctness := g.checkLetterAtPosition(letter, i)

		if correctness == wrongPosition {
			// remove the left-most occurrence
			g.positions[letter] = g.positions[letter][1:]
		}

		fb[i] = correctness
	}

	return fb
}

// markLetterAsSeen removes one occurrence of the letter from the positions map.
func (g *Gordle) markLetterAsSeen(letter rune, positionInWord int) {
	positions := g.positions[letter]

	if len(positions) == 0 {
		g.positions[letter] = nil
	}

	for i, pos := range positions {
		if pos == positionInWord {
			// remove the seen letter from the list
			g.positions[letter] = append(positions[:i], positions[i+1:]...)
			// we found it
			return
		}
	}
}

// checkLetterAtPosition returns the correctness of a letter at the specified index in the solution.
func (g *Gordle) checkLetterAtPosition(letter rune, index int) status {
	positions, ok := g.positions[letter]
	if !ok || len(positions) == 0 {
		return absentCharacter
	}

	for _, pos := range positions {
		if pos == index {
			return correctPosition
		}
	}

	return wrongPosition
}
