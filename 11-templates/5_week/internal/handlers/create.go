package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"learngo-pockets/templates/internal/habit"
)

// create takes a JSON request and creates a Habit from it,
// then redirects to index.
func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	habitName := r.FormValue("habitName")
	weeklyFreq, err := strconv.ParseInt(r.FormValue("habitFrequency"), 0, 8)
	if err != nil {
		s.logAndHideError(w, err, http.StatusBadRequest)
		return
	}

	const minFreq, maxFreq = 1, 100
	if weeklyFreq < minFreq || maxFreq < weeklyFreq {
		s.logAndHideError(w, fmt.Errorf("invalid frequency, out of bounds"), http.StatusBadRequest)
		return
	}

	err = s.client.CreateHabit(r.Context(), habit.Habit{
		Name:            habit.Name(habitName),
		WeeklyFrequency: habit.TickCount(weeklyFreq),
	})
	if err != nil {
		s.logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	// redirect to index
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
