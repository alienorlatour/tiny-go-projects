package handlers

import (
	"net/http"
	"slices"
)

var supportedAssets = []string{
	"styles.css",
}

// assets serves some identified static files. See the list above.
func (s *Server) assets(w http.ResponseWriter, r *http.Request) {
	const (
		fileNamePathValue = "habitID"
	)

	fileName := r.PathValue(fileNamePathValue)
	if fileName == "" {
		http.Error(w, "missing the name of the file", http.StatusNotFound)
		return
	}

	// prevent injection
	if !isValidAsset(fileName) {
		http.Error(w, "file not found", http.StatusNotFound)
	}

	http.ServeFile(w, r, "internal/assets/"+fileName)
}

func isValidAsset(name string) bool {
	return slices.Contains(supportedAssets, name)
}
