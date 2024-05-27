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
	const indexEndpoint = "index"

	habits, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		s.logAndHideError(w, indexEndpoint, err, http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		s.logAndHideError(w, indexEndpoint, err, http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, habits)
	if err != nil {
		s.lgr.Logf("Error in %s: %s", indexEndpoint, err.Error())
		// Calling http.Error here would have no effect, as we've already written the header to the writer.
		return
	}
}
