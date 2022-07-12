package gordle

import (
	"fmt"
)

const (
	wordLength = 5
)

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct{}

// New returns a Gordle variable, which can be used to Play!
func New() *Gordle {
	g := &Gordle{}

	return g
}

// Play runs the game.
func (g *Gordle) Play() {
	fmt.Printf("Enter a %d-letter guess:\n", wordLength)
}
