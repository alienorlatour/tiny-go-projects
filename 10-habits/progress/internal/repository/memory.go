package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/isoweek"
)

type ticksPerWeek map[isoweek.ISO8601][]time.Time

// HabitRepository holds all the current habits.
type HabitRepository struct {
	storage      map[habit.ID]habit.Habit
	ticksStorage map[habit.ID]ticksPerWeek
}

// New creates an empty habit repository.
func New() *HabitRepository {
	return &HabitRepository{
		storage:      make(map[habit.ID]habit.Habit),
		ticksStorage: make(map[habit.ID]ticksPerWeek),
	}
}

// Add inserts for the first time a habit in memory.
func (r *HabitRepository) Add(_ context.Context, habit habit.Habit) error {
	log.Print("Adding a habit...")
	r.storage[habit.ID] = habit

	return nil
}

func (r *HabitRepository) Find(_ context.Context, id habit.ID) (habit.Habit, error) {
	log.Print("Finding a habit...")
	h, found := r.storage[id]
	if !found {
		return habit.Habit{}, fmt.Errorf("habit %q not registered: %w", id, ErrNotFound)
	}

	return h, nil
}

// FindAll returns all habits.
func (r *HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	log.Printf("Listing habits...")

	habits := make([]habit.Habit, 0)
	for _, h := range r.storage {
		habits = append(habits, h)
	}

	return habits, nil
}

// AddTick inserts a new event for a habit in memory.
func (r *HabitRepository) AddTick(_ context.Context, id habit.ID, t time.Time, w isoweek.ISO8601) error {
	log.Print("Adding a tick...")
	_, ok := r.ticksStorage[id]
	if !ok {
		r.ticksStorage[id] = make(ticksPerWeek)
	}

	ticks, ok := r.ticksStorage[id][w]
	if !ok {
		ticks = make([]time.Time, 0)
	}

	r.ticksStorage[id][w] = append(ticks, t)

	return nil
}

// FindAllTicks returns all the ticks for a habit.
func (r *HabitRepository) FindAllTicks(_ context.Context, id habit.ID) ([]time.Time, error) {
	log.Printf("Listing ticks for a habit...")
	ticks := make([]time.Time, 0)
	for _, weeklyTicks := range r.ticksStorage[id] {
		ticks = append(ticks, weeklyTicks...)
	}
	return ticks, nil
}

// FindWeeklyTicks returns all the ticks in a week.
func (r *HabitRepository) FindWeeklyTicks(_ context.Context, id habit.ID, w isoweek.ISO8601) ([]time.Time, error) {
	log.Printf("Listing weekly ticks for a habit...")

	loggedWeeks, found := r.ticksStorage[id]
	if !found {
		return nil, fmt.Errorf("id %q not registered: %w", id, ErrNotFound)
	}

	if loggedWeeks[w] == nil {
		// if there is no ticks for this week, let's return an empty array instead of nil
		return []time.Time{}, nil
	}

	return loggedWeeks[w], nil
}
