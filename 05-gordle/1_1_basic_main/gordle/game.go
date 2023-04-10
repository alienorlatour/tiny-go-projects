package gordle

import (
	"fmt"
)

// Game holds all the information we need to play a game of Gordle.
type Game struct{}

// New returns a Game, which can be used to Play!
func New() *Game {
	g := &Game{}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess:\n")
}
