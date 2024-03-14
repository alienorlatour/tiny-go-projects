package handlers

import (
	"net/http"

	"learngo-pockets/templates/internal/habit"

	"github.com/go-chi/chi/v5"
)

// tick adds a tick to the given habit and redirected to index.
func (s *Server) tick(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "habitID")

	err := s.client.TickHabit(r.Context(), habit.ID(id))
	if err != nil {
		s.logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	// redirect to index
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
