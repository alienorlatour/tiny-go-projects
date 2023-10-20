package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// Server is the implementation of the grpc server.
type Server struct {
	db repository
}

type repository interface {
	Add(habit habit.Habit) error
	FindAll() ([]habit.Habit, error)
}

// New returns a Server that can Listen.
func New(repo repository) *Server {
	return &Server{
		db: repo,
	}
}

// Listen starts the listening to the port
func (s *Server) Listen(_ context.Context, port int) error {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)
	slog.Info(fmt.Sprintf("gRPC server started and listening to port %d", port))

	// Listen to the port. This will only return when something kills or stops the server.
	// TODO: Run this in a goroutine that writes to an errChan so we can select {case <- ctx.Done: ... case <- errChan: ... }
	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	// Stop or GracefulStop was called, no reason to be alarmed.
	return nil
}

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.Habit, error) {
	slog.Info(fmt.Sprintf("CreateHabit request received: %s", request))

	if request.Habit.Frequency == nil || uint(*request.Habit.Frequency) == 0 {
		return nil, fmt.Errorf("invalid frequency")
	}
	freq := *request.Habit.Frequency

	habit := habit.Habit{
		ID:        habit.ID(uuid.NewString()),
		Name:      request.Habit.Name,
		Frequency: uint(freq),
	}

	err := s.db.Add(habit)
	if err != nil {
		return nil, fmt.Errorf("cannot save habit %v: %w", habit, err)
	}

	return request.Habit, nil
}
