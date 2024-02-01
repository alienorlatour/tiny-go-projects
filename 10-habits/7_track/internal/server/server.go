package server

import (
	"context"
	"fmt"
	"net"
	_ "net/http/pprof"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
)

// Server is the implementation of the gRPC server.
type Server struct {
	api.UnimplementedHabitsServer

	lgr Logger
	db  Repository
}

// A Repository is used by the Server to interact with the database.
type Repository interface {
	Add(ctx context.Context, habit habit.Habit) error
	Find(ctx context.Context, id habit.ID) (habit.Habit, error)
	FindAll(ctx context.Context) ([]habit.Habit, error)
	AddTick(ctx context.Context, id habit.ID, t time.Time) error
	FindAllTicks(ctx context.Context, id habit.ID) ([]time.Time, error)
	FindWeeklyTicks(ctx context.Context, id habit.ID, t time.Time) ([]time.Time, error)
}

// New returns a Server that can ListenAndServe.
func New(repo Repository, lgr Logger) *Server {
	return &Server{
		db:  repo,
		lgr: lgr,
	}
}

// ListenAndServe starts listening to the port and serving requests.
func (s *Server) ListenAndServe(ctx context.Context, port int) error {
	const addr = "127.0.0.1"

	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := s.registerGRPCServer()
	s.lgr.Logf("gRPC server started and listening to port %d", port)

	// Use a channel to report errors from the gRPC server back to
	errChan := make(chan error)
	g := errgroup.Group{}
	defer func() {
		err := g.Wait()
		if err != nil {
			errChan <- fmt.Errorf("error while serving: %w", err)
		}
		close(errChan)
	}()

	// ListenAndServe to the port. This will only return when something kills or stops the server.
	g.Go(func() error {
		// This goroutine will be killed when the context is ended at the end of this function.
		err := grpcServer.Serve(listener)
		if err != nil {
			s.lgr.Logf("error while serving gRPC: %s", err)

			return fmt.Errorf("gRPC server error: %w", err)
		}

		return nil
	})

	select {
	case <-ctx.Done():
		// Stop or GracefulStop was called, no reason to be alarmed.
		s.lgr.Logf("Shutting down grpc server: %s", ctx.Err())
	case err = <-errChan:
		s.lgr.Logf("unable to serve: %w", err)
	}

	grpcServer.GracefulStop()
	_ = listener.Close()
	return nil
}

func (s *Server) registerGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerInterceptor(s.lgr)))
	api.RegisterHabitsServer(grpcServer, s)
	reflection.Register(grpcServer) // if env == dev
	return grpcServer
}

// Logger used by the server
type Logger interface {
	Logf(format string, args ...any)
}
