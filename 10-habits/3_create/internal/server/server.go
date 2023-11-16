package server

import (
	"fmt"
	"net"

	"learngo-pockets/habits/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is the implementation of the grpc server.
type Server struct {
}

// New returns a Server that can Listen.
func New() *Server {
	return &Server{}
}

// Listen starts the listening to the port
func (s *Server) Listen(port int) error {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)
	reflection.Register(grpcServer) // if env == dev

	fmt.Printf("starting server on port %d\n", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	// Stop or GracefulStop was called, no reason to be alarmed.
	return nil
}
