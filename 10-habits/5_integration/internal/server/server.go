package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strconv"

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

	errChan := make(chan error)
	// Listen to the port. This will only return when something kills or stops the server.
	go func() {
		// This goroutine will be killed when the context is ended at the end of this function.
		err := grpcServer.Serve(listener)
		if err != nil {
			errChan <- fmt.Errorf("error while listening: %w", err)
		}
	}()

	go func() {
		const pprofPort = 6060
		log.Infof("Starting pprof listener on port %d\n", pprofPort)
		err := http.ListenAndServe(net.JoinHostPort(addr, strconv.Itoa(pprofPort)), nil)
		log.Infof("error while serving pprof: %s", err)
	}()

	select {
	case <-ctx.Done():
		// Stop or GracefulStop was called, no reason to be alarmed.
		log.Infof("Shutting down grpc server: %s", ctx.Err())
		grpcServer.GracefulStop()
		return nil
	case err = <-errChan:
		grpcServer.GracefulStop()
		return fmt.Errorf("unable to serve: %w", err)
	}
}

func (s *Server) registerGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerInterceptor(s.interceptorOutput)))
	api.RegisterHabitsServer(grpcServer, s)
	reflection.Register(grpcServer) // if env == dev
	return grpcServer
}
