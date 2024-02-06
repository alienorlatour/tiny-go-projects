package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	habits, err := s.client.ListHabits(context.Background())
	if err != nil {
		http.Error(w, "Error while fetching data - please retry.", http.StatusInternalServerError)
		return
	}

	io.WriteString(w, `
<head><title>Learn Go</title></head>
<body><h2>Habits</h2>
<ul>
`)
	for _, h := range habits {
		io.WriteString(w, fmt.Sprintf("<li>%s - %d</li>", h.Name, h.WeeklyFrequency))
	}
	io.WriteString(w, `
</ul>
</body>
`)
}
