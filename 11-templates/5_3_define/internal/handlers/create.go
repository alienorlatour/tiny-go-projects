package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"learngo-pockets/templates/internal/habit"
)

// create takes a form request and creates a Habit from it,
// then redirects to index.
func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	const createEndpoint = "create"

	habitName := r.FormValue("habitName")
	weeklyFreq, err := strconv.Atoi(r.FormValue("habitFrequency"))
	if err != nil {
		s.logAndHideError(w, createEndpoint, err, http.StatusBadRequest)
		return
	}

	// limit frequency to a sensible range.
	const minFreq, maxFreq = 1, 100
	if weeklyFreq < minFreq || maxFreq < weeklyFreq {
		s.logAndHideError(w, createEndpoint, fmt.Errorf("invalid frequency, out of bounds"), http.StatusBadRequest)
		return
	}

	err = s.client.CreateHabit(r.Context(), habit.Habit{
		Name:            habit.Name(habitName),
		WeeklyFrequency: habit.TickCount(weeklyFreq),
	})
	if err != nil {
		s.logAndHideError(w, createEndpoint, err, http.StatusInternalServerError)
		return
	}

	// redirect to index
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
