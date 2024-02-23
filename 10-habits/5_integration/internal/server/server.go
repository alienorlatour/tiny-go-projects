package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/log"
)

// Server is the implementation of the gRPC server.
type Server struct {
	api.UnimplementedHabitsServer
	interceptorOutput io.Writer

	db Repository
}

type Repository interface {
	Add(ctx context.Context, habit habit.Habit) error
	FindAll(ctx context.Context) ([]habit.Habit, error)
}

// New returns a Server that can Listen.
func New(interceptorOutput io.Writer, repo Repository) *Server {
	return &Server{
		interceptorOutput: interceptorOutput,
		db:                repo,
	}
}

// Listen starts the listening to the port.
func (s *Server) Listen(ctx context.Context, port int) error {
	const addr = "127.0.0.1"

	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := s.registerGRPCServer()
	log.Infof("gRPC server started and listening to port %d", port)

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

	// Listen to the port. This will only return when something kills or stops the server.
	g.Go(func() error {
		// This goroutine will be killed when the context is ended at the end of this function.
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Infof("error while serving gRPC: %s", err)

			return fmt.Errorf("gRPC server error: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		const pprofPort = 6060
		log.Infof("Starting pprof listener on port %d\n", pprofPort)
		err := http.ListenAndServe(net.JoinHostPort(addr, strconv.Itoa(pprofPort)), nil)
		if err != nil {
			log.Infof("error while serving pprof: %s", err)

			return fmt.Errorf("pprof server error: %w", err)
		}

		return nil
	})

	select {
	case <-ctx.Done():
		// Stop or GracefulStop was called, no reason to be alarmed.
		log.Infof("Shutting down grpc server: %s", ctx.Err())
		grpcServer.GracefulStop()
		return nil
	case err = <-errChan:
		grpcServer.GracefulStop()
		if err != nil {
			return fmt.Errorf("unable to serve: %w", err)
		}

		return nil
	}
}

func (s *Server) registerGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerInterceptor(s.interceptorOutput)))
	api.RegisterHabitsServer(grpcServer, s)
	reflection.Register(grpcServer) // if env == dev
	return grpcServer
}
