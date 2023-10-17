package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"learngo-pockets/habits/api"
)

// Server is the implementation of the grpc server.
type Server struct{}

// New returns a Server that can Listen()
// TODO: A discuter: New peut etre appele sans Listen(). Je trouve qu'on devrait tout mettre dans un bloc.
// Genre StartAndListen(ctx, port) error. S'il faut se connecter a une DB, on met ca ou ?
func New() *Server {
	return &Server{}
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

func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.Habit, error) {
	slog.Info(fmt.Sprintf("CreateHabit request received: %s", request))
	// TODO implement me
	panic("implement me")
}

// Check at compilation time that we implement the grpc API.
var _ api.HabitsServer = (*Server)(nil)
