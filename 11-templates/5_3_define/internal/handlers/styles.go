package handlers

import (
	_ "embed"
	"net/http"
	"text/template"
)

//go:embed templates/styles.css
var stylesPage string

func (s *Server) styles(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("styles").Parse(stylesPage)
	if err != nil {
		s.lgr.Logf("can't parse styles.css: %s", err.Error())
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "text/css")

	err = tpl.Execute(w, map[string]interface{}{
		"Background": "DarkSlateGray",
		"Foreground": "LightGray"})
	if err != nil {
		s.lgr.Logf("cannot render styles: %s", err.Error())
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}
}
