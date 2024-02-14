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
		// log.Logger().Errorf("can't parse index: %s", err) FIXME log
		http.Error(w, "Error while rendering - please retry.", http.StatusInternalServerError)
	}

	err = tpl.Execute(w, 5)
	if err != nil {
		// log.Errorf(r.Context(), "cannot render index: %s", err) FIXME log
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}
}
