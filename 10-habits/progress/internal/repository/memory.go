package repository

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/isoweek"
)

// ticksPerWeek holds all the timestamps for a given week number.
type ticksPerWeek map[isoweek.ISO8601][]time.Time

// HabitRepository holds all the current habits.
type HabitRepository struct {
	mutex  sync.Mutex
	habits map[habit.ID]habit.Habit
	ticks  map[habit.ID]ticksPerWeek
}

// New creates an empty habit repository.
func New() *HabitRepository {
	return &HabitRepository{
		habits: make(map[habit.ID]habit.Habit),
	}
}

// Add inserts for the first time a habit in memory.
func (hr *HabitRepository) Add(_ context.Context, habit habit.Habit) error {
	log.Print("Adding a habit...")

	// Lock the writing of the habit.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	hr.habits[habit.ID] = habit

	return nil
}

func (hr *HabitRepository) Find(_ context.Context, id habit.ID) (habit.Habit, error) {
	log.Print("Finding a habit...")
	h, found := hr.habits[id]
	if !found {
		return habit.Habit{}, fmt.Errorf("habit %q not registered: %w", id, ErrNotFound)
	}

	return h, nil
}

// FindAll returns all habits.
func (hr *HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	log.Printf("Listing habits...")

	// Lock the reading and the writing of the habits.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	habits := make([]habit.Habit, 0)
	for _, h := range hr.habits {
		habits = append(habits, h)
	}

	return habits, nil
}

// AddTick inserts a new event for a habit in memory.
func (hr *HabitRepository) AddTick(_ context.Context, id habit.ID, t time.Time) error {
	log.Print("Adding a tick...")

	// Lock the reading and the writing of the habits.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	_, found := hr.ticks[id]
	if !found {
		hr.ticks[id] = make(ticksPerWeek)
	}

	w := isoweek.At(t)

	ticks, found := hr.ticks[id][w]
	if !found {
		// Capacity is set to 1 since we will add only one tick here.
		ticks = make([]time.Time, 0, 1)
	}

	hr.ticks[id][w] = append(ticks, t)

	return nil
}

// FindAllTicks returns all the ticks for a habit.
func (hr *HabitRepository) FindAllTicks(_ context.Context, id habit.ID) ([]time.Time, error) {
	log.Printf("Listing ticks for a habit...")

	// Lock the reading and the writing of the habits.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	ticks := make([]time.Time, 0)
	for _, weeklyTicks := range hr.ticks[id] {
		ticks = append(ticks, weeklyTicks...)
	}
	return ticks, nil
}

// FindWeeklyTicks returns all the ticks in a week.
func (hr *HabitRepository) FindWeeklyTicks(_ context.Context, id habit.ID, t time.Time) ([]time.Time, error) {
	log.Printf("Listing weekly ticks for a habit...")

	// Lock the reading and the writing of the habits.
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	loggedWeeks, found := hr.ticks[id]
	if !found {
		return nil, fmt.Errorf("id %q not registered: %w", id, ErrNotFound)
	}

	w := isoweek.At(t)
	if loggedWeeks[w] == nil {
		// if there is no ticks for this week, let's return an empty array instead of nil
		return []time.Time{}, nil
	}

	return loggedWeeks[w], nil
}
