package handlers

import (
	"io"
	"net/http"
)

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, `<!DOCTYPE html>
<html>
<head>
    <title>Learn Go</title>
</head>
<body>
<h1>YAY</h1>
<p>Way to go!</p>
</body>
</html>
`)

	if err != nil {
		// log.Errorf(r.Context(), "cannot render index: %s", err) FIXME log
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}
}
