package main

import (
	"fmt"
	"net/http"
	"time"
)

func newGameHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := time.Now().Unix()

	// Header should be set before the writer.Write call.
	writer.WriteHeader(http.StatusCreated)

	_, err := writer.Write([]byte(fmt.Sprintf("{\"gameID\": \"%d\"}", id)))
	if err != nil {
		http.Error(writer, "failed to write response", http.StatusInternalServerError)
	}
}
