package repository

import (
	"fmt"
	"log"
	"sync"

	"learngo-pockets/httpgordle/internal/session"
)

// GameRepository holds all the current games.
type GameRepository struct {
	mutex   sync.Mutex
	storage map[session.GameID]session.Game
}

// New creates an empty game repository.
func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

// Add inserts for the first time a game in memory.
func (gr *GameRepository) Add(game session.Game) error {
	log.Print("Adding a game...")

	// Lock the reading and the writing of the game.
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("%w (%s)", ErrConflictingID, game.ID)
	}

	gr.storage[game.ID] = game

	return nil
}

// Find a game based on its ID. If nothing is found, return a nil pointer and an ErrNotFound error.
func (gr *GameRepository) Find(id session.GameID) (session.Game, error) {
	log.Printf("Looking for game %s...", id)

	// Lock the reading of the game.
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	game, found := gr.storage[id]
	if !found {
		return session.Game{}, fmt.Errorf("can't find game %s: %w", id, ErrNotFound)
	}

	return game, nil
}

// Update a game in the database, overwriting it.
func (gr *GameRepository) Update(game session.Game) error {
	// Lock the reading and the writing of the game.
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	_, found := gr.storage[game.ID]
	if !found {
		return fmt.Errorf("can't find game %s: %w", game.ID, ErrNotFound)
	}

	gr.storage[game.ID] = game
	return nil
}
