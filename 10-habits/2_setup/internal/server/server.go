package server

import (
	"fmt"
	"net"
	"strconv"

	"learngo-pockets/habits/api"

	"google.golang.org/grpc"
)

// Server is the implementation of the gRPC server.
type Server struct {
	api.UnimplementedHabitsServer
}

// New returns a Server that can Listen.
func New() *Server {
	return &Server{}
}

// Listen starts the listening to the port
func (s *Server) Listen(port int) error {
	const addr = "127.0.0.1"

	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	fmt.Printf("starting server on port %d\n", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	// Stop or GracefulStop was called, no reason to be alarmed.
	return nil
}
