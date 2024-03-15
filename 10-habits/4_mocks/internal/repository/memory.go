package repository

import (
	"context"
	"sort"
	"sync"

	"learngo-pockets/habits/internal/habit"
)

// HabitRepository holds all the current habits.
type HabitRepository struct {
	mutex   sync.Mutex
	storage map[habit.ID]habit.Habit
	lgr     Logger
}

// New creates an empty habit repository.
func New(lgr Logger) *HabitRepository {
	return &HabitRepository{
		lgr:     lgr,
		storage: make(map[habit.ID]habit.Habit),
	}
}

// Add inserts for the first time a habit in memory.
func (hr *HabitRepository) Add(_ context.Context, habit habit.Habit) error {
	hr.lgr.Logf("Adding a habit...")

	// Lock the writing of the habit.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	hr.storage[habit.ID] = habit

	return nil
}

// FindAll returns all habits.
func (hr *HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	hr.lgr.Logf("Listing habits...")

	// Lock the reading and the writing of the habits.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	habits := make([]habit.Habit, 0)
	for _, h := range hr.storage {
		habits = append(habits, h)
	}

	// Ensure the output is deterministic by sorting the habits.
	sort.Slice(habits, func(i, j int) bool {
		return habits[i].CreationTime.Before(habits[j].CreationTime)
	})

	return habits, nil
}

// Logger used by the repository
type Logger interface {
	Logf(format string, args ...any)
}
