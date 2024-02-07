package handlers

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"

	"learngo-pockets/habits/api"
)

//go:embed index.html
var index string

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	tpl, err := template.New("index").Parse(index)
	if err != nil {
		fmt.Println("can't parse index: ", err)
		return
	}

	resp, err := s.habitClient.ListHabits(r.Context(), &api.ListHabitsRequest{})
	if err != nil {
		fmt.Println("can't call habits service: ", err)
		return
	}

	err = tpl.Execute(w, resp.Habits)

	if err != nil {
		fmt.Println("Error in index:", err)
	}
}
