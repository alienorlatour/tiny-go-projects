package repository

import (
	"context"
	"log"

	"learngo-pockets/habits/internal/habit"
)

// HabitRepository holds all the current habits.
type HabitRepository struct {
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
	log.Print("Adding a habit...")
	hr.storage[habit.ID] = habit

	return nil
}

// FindAll returns all habits.
func (hr *HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	log.Printf("Listing habits...")

	habits := make([]habit.Habit, 0)
	for _, h := range hr.storage {
		habits = append(habits, h)
	}

	return habits, nil
}
