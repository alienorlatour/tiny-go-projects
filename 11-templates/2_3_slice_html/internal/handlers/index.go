package handlers

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed index.html
var indexPage string

func (Server) index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	tpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		fmt.Println("can't parse index: ", err)
	}

	values := []int{47, 52, 88, 18}

	err = tpl.Execute(w, values)

	if err != nil {
		fmt.Println("Error in index:", err)
	}
}
