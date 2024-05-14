package handlers

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed index.html
var indexPage string

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		s.lgr.Logf("can't parse index index: %s", err.Error())
		http.Error(w, "Error while rendering - please retry.", http.StatusInternalServerError)
	}

	err = tpl.Execute(w, 5)
	if err != nil {
		s.lgr.Logf("cannot render index: %s", err.Error())
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}
}
