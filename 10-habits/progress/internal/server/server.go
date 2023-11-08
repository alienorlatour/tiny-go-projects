package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

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
func (s *Server) Listen(ctx context.Context, port int) error {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerInterceptor(os.Stdout)))
	api.RegisterHabitsServer(grpcServer, s)
	log.Printf("gRPC server started and listening to port %d", port)

	errChan := make(chan error)
	// Listen to the port. This will only return when something kills or stops the server.
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			errChan <- fmt.Errorf("error while listening: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		// Stop or GracefulStop was called, no reason to be alarmed.
		log.Printf("Shutting down grpc server: %s", ctx.Err())
		grpcServer.GracefulStop()
		return nil
	case err = <-errChan:
		return fmt.Errorf("unable to serve: %w", err)
	}
}
