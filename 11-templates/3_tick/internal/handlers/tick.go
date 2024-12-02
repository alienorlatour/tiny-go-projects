package handlers

import (
	"net/http"

	"learngo-pockets/templates/internal/habit"
)

// tick adds a tick to the given habit and redirects to index.
func (s *Server) tick(w http.ResponseWriter, r *http.Request) {
	const (
		tickEndpoint     = "tick"
		habitIDPathValue = "habitID"
	)

	id := r.PathValue(habitIDPathValue)
	if id == "" {
		http.Error(w, "missing the id of the habit", http.StatusNotFound)
		return
	}

	err := s.client.TickHabit(r.Context(), habit.ID(id))
	if err != nil {
		s.logAndHideError(w, tickEndpoint, err, http.StatusInternalServerError)
		return
	}

	// redirect to index
	http.Redirect(w, r, indexPath, http.StatusSeeOther)
}
