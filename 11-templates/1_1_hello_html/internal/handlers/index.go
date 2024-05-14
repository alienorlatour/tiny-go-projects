package handlers

import (
	"io"
	"net/http"
)

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
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
		s.lgr.Logf("cannot render index: %s", err)
		http.Error(w, "Error while rendering.", http.StatusInternalServerError)
	}
}
