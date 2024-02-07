package handlers

import (
	"fmt"
	"net/http"

	"learngo-pockets/templates/internal/habit"

	"github.com/go-chi/chi/v5"
)

// tick adds a tick to the given habit and redirected to index.
func (s *Server) tick(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Printf("Ticking habit id %s\n", id)

	err := s.client.TickHabit(r.Context(), habit.ID(id))
	if err != nil {
		logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	// cheap redirect
	s.index(w, r)
}
