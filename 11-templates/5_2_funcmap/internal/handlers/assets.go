package handlers

import (
	_ "embed"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

// assets serves some identified files. See the list above.
func (s *Server) assets(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "filename")

	// not really a good approach TODO
	f, ok := map[string]http.HandlerFunc{
		"styles.css": s.styles,
	}[fileName]

	// prevent injection
	if !ok {
		http.Error(w, "file not found", http.StatusNotFound)
	}

	if f == nil {
		http.ServeFile(w, r, "internal/handlers/assets/static/"+fileName)
	}

	s.lgr.Logf("generating file %s", fileName)
	f(w, r)
}

//go:embed assets/dynamic/styles.css
var stylesPage string

func (s *Server) styles(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("styles").Parse(stylesPage)
	if err != nil {
		s.lgr.Logf("can't parse styles.css: %s", err.Error())
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "text/css")

	err = tpl.Execute(w, "#023047")
	if err != nil {
		s.lgr.Logf("cannot render styles: %s", err.Error())
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}
}
