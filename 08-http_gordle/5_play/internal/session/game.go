package session

import "learngo-pockets/httpgordle/internal/gordle"

// A GameID represents the ID of a game
type GameID string

// Game contains the information about a game.
type Game struct {
	// ID is the identified of a game.
	ID GameID

	// The game of Gordle that is being played.
	Gordle gordle.Game

	// AttemptsLeft counts the number of attempts left before the game is over.
	AttemptsLeft byte

	// Guesses is the list of past guesses, and their feedback.
	Guesses []Guess

	// Status tells whether the game is playable.
	Status Status
}

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
