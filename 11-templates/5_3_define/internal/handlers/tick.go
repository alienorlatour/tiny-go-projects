package handlers

import (
	"net/http"

	"learngo-pockets/templates/internal/habit"

	"github.com/go-chi/chi/v5"
)

// tick adds a tick to the given habit and redirects to index.
func (s *Server) tick(w http.ResponseWriter, r *http.Request) {
	const tickEndpoint = "tick"

	id := chi.URLParam(r, "habitID")

	err := s.client.TickHabit(r.Context(), habit.ID(id))
	if err != nil {
		s.logAndHideError(w, tickEndpoint, err, http.StatusInternalServerError)
		return
	}

	// redirect to index
	http.Redirect(w, r, indexPath, http.StatusSeeOther)
}
