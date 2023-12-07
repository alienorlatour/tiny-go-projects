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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"learngo-pockets/habits/api"
)

const port = 28710

func TestIntegration(t *testing.T) {
	// run server
	grpcServ := newServer()
	go startServer(t, grpcServ)
	defer grpcServ.Stop()

	// create client
	habitsCli, err := newClient()
	require.NoError(t, err)

	// add 2 habits
	addHabit(t, habitsCli, nil, "walk in the forest")

	addHabit(t, habitsCli, ptr(3), "read a few pages")

	addHabitWithError(t, habitsCli, 5, "  	  ", codes.InvalidArgument)

	// check that the 2 habits are present
	list := listHabits(t, habitsCli)
	assert.ElementsMatch(t, list.Habits, []*api.Habit{
		{
			Name:            "walk in the forest",
			WeeklyFrequency: 1,
		},
		{
			Name:            "read a few pages",
			WeeklyFrequency: 3,
		},
	})
}

func addHabitWithError(t *testing.T, habitsCli api.HabitsClient, freq int32, name string, statusCode codes.Code) {
	t.Helper()

	_, err := habitsCli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            name,
		WeeklyFrequency: &freq,
	})
	statusErr, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, statusCode, statusErr.Code())
}

func listHabits(t *testing.T, habitsCli api.HabitsClient) *api.ListHabitsResponse {
	list, err := habitsCli.ListHabits(context.Background(), &api.ListHabitsRequest{})
	assert.NoError(t, err)
	return list
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

func addHabit(t *testing.T, habitsCli api.HabitsClient, freq *int32, name string) {
	t.Helper()

	_, err := habitsCli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            name,
		WeeklyFrequency: freq,
	})
	assert.NoError(t, err)
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
