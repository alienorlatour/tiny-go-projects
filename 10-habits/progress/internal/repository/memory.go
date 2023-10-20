package repository

import (
	"log"
	"sync"

	"learngo-pockets/habits/internal/habit"
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
func (gr *HabitRepository) Add(habit habit.Habit) error {
	log.Print("Adding a habit...")
	gr.storage[habit.ID] = habit

	return nil
}

// FindAll returns all habits.
func (gr *HabitRepository) FindAll() ([]habit.Habit, error) {
	log.Printf("Listing habits...")

	habits := make([]habit.Habit, 0)
	for _, h := range gr.storage {
		habits = append(habits, h)
	}

	return habits, nil
}
