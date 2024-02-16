package repository

import (
	"context"
	"sync"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/log"
)

// HabitRepository holds all the current habits.
type HabitRepository struct {
	mutex   sync.Mutex
	storage map[habit.ID]habit.Habit
}

// New creates an empty habit repository.
func New() *HabitRepository {
	return &HabitRepository{
		storage: make(map[habit.ID]habit.Habit),
	}
}

// Add inserts for the first time a habit in memory.
func (hr *HabitRepository) Add(_ context.Context, habit habit.Habit) error {
	log.Infof("Adding a habit...")

	// Lock the writing of the habit.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	hr.storage[habit.ID] = habit

	return nil
}

// FindAll returns all habits.
func (hr *HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	log.Infof("Listing habits...")

	// Lock the reading and the writing of the habits.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	habits := make([]habit.Habit, 0)
	for _, h := range hr.storage {
		habits = append(habits, h)
	}

	return habits, nil
}
