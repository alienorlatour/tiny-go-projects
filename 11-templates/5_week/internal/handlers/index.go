package handlers

import (
	_ "embed"
	"html/template"
	"net/http"
	"time"

	"learngo-pockets/templates/internal/habit"
)

//go:embed index.html
var indexPage string

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	habits, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		s.logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		s.logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	week := habit.NewWeek(time.Now(), "02 January 2006")

	err = tpl.Execute(w, map[string]interface{}{
		"Habits": habits,
		"Date":   week,
	})
	if err != nil {
		s.logAndHideError(w, err, http.StatusInternalServerError)
		return
	}
}
