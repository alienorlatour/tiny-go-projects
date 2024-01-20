package repository

import (
	"context"
	"log"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/tick"
)

// TickRepository holds all the current habits with their associated events.
type TickRepository struct {
	storage map[habit.ID][]tick.Tick
}

// NewTickRepository creates an empty tick repository.
func NewTickRepository() *TickRepository {
	return &TickRepository{
		storage: make(map[habit.ID][]tick.Tick),
	}
}

// Add inserts a new event for a habit in memory.
func (tr *TickRepository) Add(_ context.Context, id habit.ID, t tick.Tick) error {
	log.Print("Adding a tick...")
	ticks := tr.storage[id]
	tr.storage[id] = append(ticks, t)

	return nil
}

// FindAll returns all the ticks for a habit.
func (tr *TickRepository) FindAll(_ context.Context, id habit.ID) ([]tick.Tick, error) {
	log.Printf("Listing ticks for a habit...")

	return tr.storage[id], nil
}
