package gordle

import (
	"fmt"
)

// Game holds all the information we need to play a game of gordle.
type Game struct{}

// New returns a Game variable, which can be used to Play!
func New() *Game {
	g := &Game{}
	fmt.Println("Welcome to Gordle!")

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Printf("Enter a guess:\n")
}
