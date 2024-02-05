package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"learngo-pockets/habits/api"
)

type Server struct {
	router chi.Router

	habitClient api.HabitsClient
}

func New(ctx context.Context, url string) *Server {
	s := &Server{}

	client, err := grpc.DialContext(ctx, url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("unable to connect to backend: %s", err.Error()))
		return nil
	}

	s.habitClient = api.NewHabitsClient(client)

	return s
}

func (s *Server) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.index)
	r.Post("/habits/{id}", s.tick)

	return r
}
