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
	_, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		fmt.Println("error!", err.Error())
		http.Error(w, "Error while fetching data - please retry.", http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		fmt.Println("can't parse index: ", err)
		http.Error(w, "Error while rendering - please retry.", http.StatusInternalServerError)
	}

	values := []int{47, 52, 88, 18}

	err = tpl.Execute(w, values)
	if err != nil {
		fmt.Println("Error in index:", err)
		http.Error(w, "Error while rendering - please retry.", http.StatusInternalServerError)
	}
}
