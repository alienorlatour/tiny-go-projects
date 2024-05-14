package handlers

import (
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
)

var supportedAssets = []string{
	"styles.css",
}

// assets serves some identified static files. See the list above.
func (s *Server) assets(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "filename")

	// prevent injection
	if !isValidAsset(fileName) {
		http.Error(w, "file not found", http.StatusNotFound)
	}

	http.ServeFile(w, r, "internal/assets/"+fileName)
}

func isValidAsset(name string) bool {
	return slices.Contains(supportedAssets, name)
}
