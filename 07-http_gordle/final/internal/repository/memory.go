package repository

import (
	"fmt"
	"learngo-pockets/httpgordle/internal/domain"
	"log"
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

// Add inserts for the first time a game in memory.
func (gr *GameRepository) Add(game domain.Game) error {
	log.Print("Adding a game...")

	_, ok := gr.games[game.ID]
	if ok {
		return fmt.Errorf("gameID %s already exists", game.ID)
	}

	gr.games[game.ID] = game

	return nil
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
