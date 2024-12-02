package handlers

import (
	"net/http"
)

// Server serves all the HTML routes on this service.
type Server struct {
	lgr Logger
}

// New builds a new server.
func New(lgr Logger) *Server {
	return &Server{lgr: lgr}
}

// Router returns an HTTP handler that listens to all the proper paths.
func (s *Server) Router() http.Handler {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodGet+" "+"/", s.index)

	return r
}

type Logger interface {
	Logf(format string, args ...any)
}
