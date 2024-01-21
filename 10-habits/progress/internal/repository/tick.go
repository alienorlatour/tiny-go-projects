package repository

import (
	"context"
	"log"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/tick"
)

// ticksPerWeek stores all the ticks per number of weeks.
type ticksPerWeek map[tick.ISOWeek][]tick.Tick

// TickRepository holds all the current habits with their associated events.
type TickRepository struct {
	storage map[habit.ID]ticksPerWeek
}

// NewTickRepository creates an empty tick repository.
func NewTickRepository() *TickRepository {

	return &TickRepository{
		storage: make(map[habit.ID]ticksPerWeek),
	}
}

// Add inserts a new event for a habit in memory.
func (tr *TickRepository) Add(_ context.Context, id habit.ID, t tick.Tick, w tick.ISOWeek) error {
	log.Print("Adding a tick...")
	_, ok := tr.storage[id]
	if !ok {
		tr.storage[id] = make(ticksPerWeek)
	}

	ticks, ok := tr.storage[id][w]
	if !ok {
		ticks = make([]tick.Tick, 0)
	}

	tr.storage[id][w] = append(ticks, t)

	return nil
}

// FindAll returns all the ticks for a habit.
func (tr *TickRepository) FindAll(_ context.Context, id habit.ID) ([]tick.Tick, error) {
	log.Printf("Listing ticks for a habit...")
	ticks := make([]tick.Tick, 0)
	for _, weeklyTicks := range tr.storage[id] {
		ticks = append(ticks, weeklyTicks...)
	}
	return ticks, nil
}

// FindWeeklyTicks returns all the ticks in a week
func (tr *TickRepository) FindWeeklyTicks(_ context.Context, id habit.ID, w tick.ISOWeek) ([]tick.Tick, error) {
	log.Printf("Listing weekly ticks for a habit...")

	return tr.storage[id][w], nil
}
