package repository

import (
	"fmt"
	"log"

	"learngo-pockets/httpgordle/internal/session"
)

// GameRepository holds all the current games.
type GameRepository struct {
	// games stores the list of games and makes them accessible with their ID.
	// TODO: Document: We could add an extra layer on top of session.Game
	games map[session.GameID]session.Game
}

// New creates an empty game repository.
func New() *GameRepository {
	return &GameRepository{
		games: make(map[session.GameID]session.Game),
	}
}

// Add inserts for the first time a game in memory.
func (gr *GameRepository) Add(game session.Game) error {
	log.Print("Adding a game...")

	_, ok := gr.games[game.ID]
	if ok {
		return fmt.Errorf("gameID %s already exists", game.ID)
	}

	gr.games[game.ID] = game

	return nil
}

// Find a game based on its ID. If nothing is found, return a nil pointer and an ErrNotFound error.
func (gr *GameRepository) Find(id session.GameID) (session.Game, error) {
	log.Printf("Looking for game %s...", id)

	game, found := gr.games[id]
	if !found {
		return session.Game{}, fmt.Errorf("can't find game %s: %w", id, ErrNotFound)
	}

	return game, nil
}

// Update a game in the database, overwriting it.
func (gr *GameRepository) Update(id session.GameID, game session.Game) error {
	_, found := gr.games[id]
	if !found {
		return fmt.Errorf("can't find game %s: %w", id, ErrNotFound)
	}

	gr.games[id] = game
	return nil
}
