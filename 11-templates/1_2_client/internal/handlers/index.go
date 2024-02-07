package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	// TODO get time from parameters
	habits, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		fmt.Println("error!", err.Error())
		http.Error(w, "Error while fetching data - please retry.", http.StatusInternalServerError)
		return
	}

	io.WriteString(w, `
<head><title>Learn Go</title></head>
<body><h2>Habits</h2>
<ul>
`)
	for _, h := range habits {
		io.WriteString(w, fmt.Sprintf("<li>%s - %d/%d</li>", h.Name, h.Ticks, h.WeeklyFrequency))
	}
	io.WriteString(w, `
</ul>
</body>
`)
}
