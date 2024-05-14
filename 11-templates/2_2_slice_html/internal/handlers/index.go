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
	_, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		s.lgr.Logf("error! %s", err.Error())
		http.Error(w, "Error while fetching data - please retry.", http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		s.lgr.Logf("can't parse index: %s", err.Error())
		http.Error(w, "Error while rendering - please retry.", http.StatusInternalServerError)
		return
	}

	values := []int{47, 52, 88, 18}

	err = tpl.Execute(w, values)
	if err != nil {
		s.lgr.Logf("Error in index: %s", err.Error())
		http.Error(w, "Error while rendering - please retry.", http.StatusInternalServerError)
	}
}
