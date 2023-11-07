package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// Server is the implementation of the grpc server.
type Server struct {
	db repository
}

type repository interface {
	Add(ctx context.Context, habit habit.Habit) error
	FindAll(ctx context.Context) ([]habit.Habit, error)
}

// New returns a Server that can Listen.
func New(repo repository) *Server {
	return &Server{
		db: repo,
	}
}

// Listen starts the listening to the port
func (s *Server) Listen(port int) error {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerIntercept))
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
