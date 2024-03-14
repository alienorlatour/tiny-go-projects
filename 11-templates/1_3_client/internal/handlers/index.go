package handlers

import (
	_ "embed"
	"html/template"
	"net/http"
	"time"
)

//go:embed index.html
var indexPage string

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	// TODO get time from parameters
	habits, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		s.lgr.Logf("error! %s", err.Error())
		http.Error(w, "Error while fetching data - please retry.", http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		s.lgr.Logf("can't parse index index: %s", err.Error())
		http.Error(w, "Error while rendering - please retry.", http.StatusInternalServerError)
	}

	err = tpl.Execute(w, len(habits))
	if err != nil {
		s.lgr.Logf("cannot render index: %s", err.Error())
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}
}
