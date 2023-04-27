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
func (gr *GameRepository) Find(id domain.GameID) *domain.Game {
	log.Printf("Looking for game %s...", id)

	return nil
}
