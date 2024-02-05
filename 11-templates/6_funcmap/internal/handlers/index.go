package handlers

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"learngo-pockets/habits/api"
)

//go:embed index.html
var index string

func scoreStatus(habit *api.Habit) string {
	res := float32(habit.WeeklyFrequency) / 5.0
	switch {
	case res == 0:
		return "not_started"
	case res < 1:
		return "started"
	case res < 5:
		return "good_progress"
	default:
		return "completed"
	}
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	tpl, err := template.New("index").Funcs(template.FuncMap{"scoreStatus": scoreStatus}).Parse(index)
	if err != nil {
		fmt.Println("can't parse index: ", err)
		return
	}

	resp, err := s.habitClient.ListHabits(r.Context(), &api.ListHabitsRequest{})
	if err != nil {
		fmt.Println("can't call habits service: ", err)
		return
	}

	var iw isoWeek
	iw.year, iw.weekNumber = time.Now().ISOWeek()

	err = tpl.Execute(w, map[string]interface{}{
		"Habits": resp.Habits,
		"Date":   iw,
	})

	if err != nil {
		fmt.Println("Error in index:", err)
	}
}

type isoWeek struct {
	weekNumber int
	year       int
}

func (iw isoWeek) Start() string {
	return time.Date(iw.year, 1, 1, 0, 0, 0, 0, time.UTC).
		AddDate(0, 0, (iw.weekNumber-1)*7).
		Format("02 January 2006")
}

func (iw isoWeek) End() string {
	return time.Date(iw.year, 1, 1, 0, 0, 0, 0, time.UTC).
		AddDate(0, 0, iw.weekNumber*7-1).
		Format("02 January 2006")
}
