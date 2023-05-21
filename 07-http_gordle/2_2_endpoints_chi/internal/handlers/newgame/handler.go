package newgame

import (
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Creating a new game"))
}
