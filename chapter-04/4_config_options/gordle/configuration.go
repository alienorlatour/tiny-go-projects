package gordle

import (
	"bufio"
)

// ConfigFunc defines a configuration function for the Gordle.
type ConfigFunc func(g *Gordle) error

// WithReader provider a specific reader from which the player suggestions will be read. Default value is stdin.
func WithReader(reader *bufio.Reader) ConfigFunc {
	return func(g *Gordle) error {
		g.reader = reader
		return nil
	}
}

// WithSolution sets the solution to the current game (default is random from corpus).
func WithSolution(solution []rune) ConfigFunc {
	return func(g *Gordle) error {
		g.solution = solution
		return nil
	}
}

// WithMaxAttempts changes the maximum number of attempts (default is unlimited)
func WithMaxAttempts(maxAttempts int) ConfigFunc {
	return func(g *Gordle) error {
		g.maxAttempts = maxAttempts
		return nil
	}
}
