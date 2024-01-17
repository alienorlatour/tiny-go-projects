package repository

import (
	"context"
	"log"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/tick"
)

// TickRepository holds all the current habits with their associated events.
type TickRepository struct {
	storage map[habit.ID]tick.Tick
}

// NewTickRepository creates an empty tick repository.
func NewTickRepository() *TickRepository {
	return &TickRepository{
		storage: make(map[habit.ID]tick.Tick),
	}
}

// Add inserts a new event for a habit in memory.
func (tr *TickRepository) Add(_ context.Context, t tick.Tick) error {
	log.Print("Adding a tick...")
	tr.storage[t.HabitID] = t

	return nil
}

// Find returns all the ticks for a habit.
func (tr *TickRepository) Find(_ context.Context) ([]tick.Tick, error) {
	log.Printf("Listing ticks for a habit...")

	ticks := make([]tick.Tick, 0)
	for _, h := range tr.storage {
		ticks = append(ticks, h)
	}

	return ticks, nil
}
