package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// Server is the implementation of the gRPC server.
type Server struct {
	api.UnimplementedHabitsServer
	db repository
}

type repository interface {
	Add(ctx context.Context, habit habit.Habit) error
	Find(ctx context.Context, id habit.ID) (habit.Habit, error)
	FindAll(ctx context.Context) ([]habit.Habit, error)
	AddTick(ctx context.Context, id habit.ID, t time.Time) error
	FindAllTicks(ctx context.Context, id habit.ID) ([]time.Time, error)
	FindWeeklyTicks(ctx context.Context, id habit.ID, t time.Time) ([]time.Time, error)
}

// New returns a Server that can Listen.
func New(repo repository) *Server {
	return &Server{
		db: repo,
	}
}

// Listen starts the listening to the port.
func (s *Server) Listen(ctx context.Context, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerInterceptor(os.Stdout)))
	api.RegisterHabitsServer(grpcServer, s)
	reflection.Register(grpcServer) // if env == dev
	log.Printf("gRPC server started and listening to port %d", port)

	errChan := make(chan error)
	// Listen to the port. This will only return when something kills or stops the server.
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			errChan <- fmt.Errorf("error while listening: %w", err)
		}
	}()

	go func() {
		const pprofPort = 6060
		log.Printf("Starting pprof listener on port %d\n", pprofPort)
		err := http.ListenAndServe(fmt.Sprintf(":%d", pprofPort), nil)
		log.Printf("error while serving pprof: %s", err)
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
