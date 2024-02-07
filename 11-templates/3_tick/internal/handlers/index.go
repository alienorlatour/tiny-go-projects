package handlers

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//go:embed index.html
var indexPage string

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	habits, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		logAndHideError(w, err, http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, habits)
	if err != nil {
		logAndHideError(w, err, http.StatusInternalServerError)
		return
	}
}

func logAndHideError(w http.ResponseWriter, err error, httpStatus int) {
	fmt.Println("Error in index:", err)
	http.Error(w, "Error while rendering - please retry.", httpStatus)
}
