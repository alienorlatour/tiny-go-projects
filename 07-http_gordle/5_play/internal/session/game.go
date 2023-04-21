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

// ErrGameOver is returned when a play is made but the game is over.
var ErrGameOver = fmt.Errorf("game over")

func NewGame(corpus []string) Game {
	game, err := gordle.New(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to create a new gordle game")
	}

	id := session.GameID(fmt.Sprintf("%d", rand.Int()))
	g := session.Game{
		ID:           id,
		Gordle:       *game,
		AttemptsLeft: maxAttempts,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
	}
}
