package handlers

import (
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
		logAndHideError(w, err, http.StatusBadRequest)
		return
	}

	err = s.client.CreateHabit(r.Context(), habit.Habit{
		Name:            habit.Name(habitName),
		WeeklyFrequency: habit.WeeklyFrequency(weeklyFreq),
	})
	if err != nil {
		logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	// redirect to index
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
