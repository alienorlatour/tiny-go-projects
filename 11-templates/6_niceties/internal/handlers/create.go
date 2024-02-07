package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"learngo-pockets/habits/api"
)

type createRequest struct {
	Name            string `json:"name"`
	WeeklyFrequency int32  `json:"weekly_frequency"`
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request to create a habit")

	cr := createRequest{}

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&cr)
	if err != nil {
		fmt.Println("can't decode request: ", err)
		return
	}

	fmt.Printf("creating habit %v\n", cr)

	_, err = s.habitClient.CreateHabit(r.Context(), &api.CreateHabitRequest{Name: cr.Name, WeeklyFrequency: &cr.WeeklyFrequency})
	if err != nil {
		fmt.Println("can't create habit: ", err)
		return
	}
}
