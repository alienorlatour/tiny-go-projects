package integration

import (
	"context"
	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"learngo-pockets/habits/api"
)

func TestIntegration(t *testing.T) {
	// run server
	grpcServ := newServer(t)
	listener, err := net.Listen("tcp", "")
	require.NoError(t, err)

	go func() {
		err = grpcServ.Serve(listener)
		require.NoError(t, err)
	}()
	defer grpcServ.Stop()

	// create client
	habitsCli, err := newClient(listener.Addr().String())
	require.NoError(t, err)

	// add 2 habits
	idWalk := addHabit(t, habitsCli, nil, "walk in the forest")

	idRead := addHabit(t, habitsCli, ptr(3), "read a few pages")

	addHabitWithError(t, habitsCli, 5, "  	  ", codes.InvalidArgument)

	// check that the 2 habits are present
	listHabitsMatches(t, habitsCli, []*api.Habit{
		{
			Name:            "walk in the forest",
			WeeklyFrequency: 1,
		},
		{
			Name:            "read a few pages",
			WeeklyFrequency: 3,
		},
	})

	// add 2 ticks for Walk habit
	tickHabit(t, habitsCli, idWalk)
	tickHabit(t, habitsCli, idWalk)

	// add 1 tick for Read habit
	tickHabit(t, habitsCli, idRead)

	getHabitStatusMatches(t, habitsCli, idWalk, &api.GetHabitStatusResponse{
		Habit: &api.Habit{
			Id:              idWalk,
			Name:            "walk in the forest",
			WeeklyFrequency: 1,
		},
		TicksCount: 2,
	})

	getHabitStatusMatches(t, habitsCli, idRead, &api.GetHabitStatusResponse{
		Habit: &api.Habit{
			Id:              idRead,
			Name:            "read a few pages",
			WeeklyFrequency: 3,
		},
		TicksCount: 1,
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

func listHabitsMatches(t *testing.T, habitsCli api.HabitsClient, expected []*api.Habit) {
	list, err := habitsCli.ListHabits(context.Background(), &api.ListHabitsRequest{})
	require.NoError(t, err)

	for i := range list.Habits {
		assert.NotEqual(t, "", list.Habits[i].Id)
		list.Habits[i].Id = "" // generated
	}
	assert.ElementsMatch(t, list.Habits, expected)
}

func tickHabit(t *testing.T, habitsCli api.HabitsClient, id string) {
	t.Helper()

	_, err := habitsCli.TickHabit(context.Background(), &api.TickHabitRequest{
		HabitId: id,
	})
	assert.NoError(t, err)
}

func getHabitStatusMatches(t *testing.T, habitsCli api.HabitsClient, id string, expected *api.GetHabitStatusResponse) {
	h, err := habitsCli.GetHabitStatus(context.Background(), &api.GetHabitStatusRequest{HabitId: id})
	require.NoError(t, err)

	assert.Equal(t, expected.Habit, h.Habit)
	assert.Equal(t, expected.TicksCount, h.TicksCount)
}

func newServer(t *testing.T) *grpc.Server {
	t.Helper()
	s := server.New(repository.New())

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	return grpcServer
}

func ptr(i int32) *int32 {
	return &i
}

func addHabit(t *testing.T, habitsCli api.HabitsClient, freq *int32, name string) string {
	t.Helper()

	resp, err := habitsCli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            name,
		WeeklyFrequency: freq,
	})
	assert.NoError(t, err)

	return resp.Habit.Id
}

func newClient(serverAddress string) (api.HabitsClient, error) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddress, creds)
	if err != nil {
		return nil, err
	}

	return api.NewHabitsClient(conn), nil
}
