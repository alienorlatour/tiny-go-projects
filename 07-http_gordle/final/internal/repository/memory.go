package repository

import (
	"fmt"
	"log"
	"math/rand"

	"learngo-pockets/httpgordle/internal/domain"
)

// GameRepository holds all the current games.
type GameRepository struct {
	// games stores the list of games and makes them accessible with their ID.
	// TODO: Document: We could add an extra layer on top of domain.Game
	games map[domain.GameID]domain.Game
}

// New creates an empty game repository.
func New() *GameRepository {
	return &GameRepository{
		games: make(map[domain.GameID]domain.Game),
	}
}

// Create a game.
func (gr *GameRepository) Create() domain.Game {
	log.Print("Creating a game...")
	return domain.Game{
		ID: domain.GameID(fmt.Sprintf("%d", rand.Int())),
	}
}

// Find a game based on its ID. If nothing is found, return a nil pointer.
func (gr *GameRepository) Find(id domain.GameID) (domain.Game, error) {
	log.Printf("Looking for game %s...", id)

	game, found := gr.games[id]
	if !found {
		return domain.Game{}, fmt.Errorf("can't find game %s: %w", id, ErrNotFound)
	}

	return game, nil
}

// Guess tries to guess if the word is the solution.
// This returns the new state of the game, including the feedback for this guess,
// or an error, if the guess was invalid.
func (gr *GameRepository) Guess(id domain.GameID, guess string) (domain.Game, error) {
	game, found := gr.games[id]
	if !found {
		return domain.Game{}, fmt.Errorf("can't guess in game %s: %w", id, ErrNotFound)
	}

	// TODO: Guess.

	return game, nil
}
