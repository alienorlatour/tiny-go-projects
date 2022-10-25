package gordle

import (
	"bufio"
	"io"
)

// ConfigFunc defines a configuration function for the Game.
type ConfigFunc func(g *Game) error

// WithReader provider a specific reader from which the player suggestions will be read. Default value is stdin.
func WithReader(reader io.Reader) ConfigFunc {
	return func(g *Game) error {
		g.reader = bufio.NewReader(reader)
		return nil
	}
}

// WithSolution sets the solution to the current game (default is random from corpus).
func WithSolution(solution []rune) ConfigFunc {
	return func(g *Game) error {
		g.solution = solution
		return nil
	}
}

// WithMaxAttempts changes the maximum number of attempts (default is unlimited)
func WithMaxAttempts(maxAttempts int) ConfigFunc {
	return func(g *Game) error {
		g.maxAttempts = maxAttempts
		return nil
	}
}
