package session

import (
	"fmt"

	"learngo-pockets/httpgordle/internal/gordle"
)

// Game contains the information about a game.
type Game struct {
	ID           GameID
	Gordle       gordle.Game
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

// A GameID represents the ID of a game.
type GameID string

// Status is the current status of the game and tells what operations can be made on it.
type Status string

const (
	StatusPlaying = "Playing"
	StatusWon     = "Won"
	StatusLost    = "Lost"
)

// A Guess is a pair of a word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string
	Feedback string
}

// ErrGameOver can signal when the game cannot be played any further because it is over.
var ErrGameOver = fmt.Errorf("game over")
