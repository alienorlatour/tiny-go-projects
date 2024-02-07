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

func (Server) index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		fmt.Println("can't parse index: ", err)
	}

	err = tpl.Execute(w, time.Now().Format(time.RFC3339))

	if err != nil {
		fmt.Println("Error in index:", err)
	}
}
