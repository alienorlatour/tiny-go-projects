package integration

import (
	"context"
	"fmt"
	"net"
	"testing"

	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"learngo-pockets/habits/api"
)

const port = 28710

func TestIntegration(t *testing.T) {
	ctx := context.Background()

	grpcServ := newServer()
	go startServer(t, grpcServ)
	defer grpcServ.Stop()

	habitsCli, err := newClient()
	require.NoError(t, err)

	_, err = addHabit(habitsCli, ctx, 5, "walk in the forest")
	require.NoError(t, err)

	_, err = addHabit(habitsCli, ctx, 3, "read a few pages")
	require.NoError(t, err)

	list, err := habitsCli.ListHabits(ctx, &api.ListHabitsRequest{})
	assert.NoError(t, err)
	assert.ElementsMatch(t, list.Habits, []*api.Habit{
		{
			Name:      "walk in the forest",
			Frequency: ptr(5),
		},
		{
			Name:      "read a few pages",
			Frequency: ptr(3),
		},
	})
}

func startServer(t *testing.T, grpcServ *grpc.Server) {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	require.NoError(t, err)

	err = grpcServ.Serve(listener)
	require.NoError(t, err)
}

func newServer() *grpc.Server {
	s := server.New(repository.New())

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	return grpcServer
}

func ptr(i int32) *int32 {
	return &i
}

func addHabit(habitsCli api.HabitsClient, ctx context.Context, walkFrequency int32, name string) (*api.CreateHabitResponse, error) {
	return habitsCli.CreateHabit(ctx, &api.CreateHabitRequest{Habit: &api.Habit{Name: name, Frequency: &walkFrequency}})
}

func newClient() (api.HabitsClient, error) {
	serverAddress := fmt.Sprintf(":%d", port)
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	habitsCli := api.NewHabitsClient(conn)
	return habitsCli, nil
}
