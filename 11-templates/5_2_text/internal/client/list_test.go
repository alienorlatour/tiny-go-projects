package client_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"learngo-pockets/habits/api"
	"learngo-pockets/templates/internal/client"
	"learngo-pockets/templates/internal/client/mocks"
	"learngo-pockets/templates/internal/habit"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

//go:generate minimock -i learngo-pockets/habits/api.HabitsClient -s "_mock.go" -o "mocks"

func TestListHabits(t *testing.T) {
	now := time.Now()

	// Create a mock for the API client
	mockClient := mocks.NewHabitsClientMock(t)

	// Create a HabitsClient with the mock client
	habitsClient := client.New(mockClient)

	// Define sample data for the mock response
	mockResponse := &api.ListHabitsResponse{
		Habits: []*api.Habit{
			{Id: "ID1", Name: "Knit", WeeklyFrequency: 3},
			{Id: "ID2", Name: "Code", WeeklyFrequency: 5},
		},
	}

	mockClient.ListHabitsMock.Expect(minimock.AnyContext, &api.ListHabitsRequest{}).Return(mockResponse, nil)
	mockClient.GetHabitStatusMock.Set(func(_ context.Context, in *api.GetHabitStatusRequest, _ ...grpc.CallOption) (*api.GetHabitStatusResponse, error) {
		if in.HabitId == "ID1" {
			return &api.GetHabitStatusResponse{TicksCount: 3}, nil
		}
		if in.HabitId == "ID2" {
			return &api.GetHabitStatusResponse{TicksCount: 2}, nil
		}
		return nil, fmt.Errorf("unexpected ID")
	})

	// Call the function being tested
	habits, err := habitsClient.ListHabits(context.Background(), now)

	// Assert that there are no errors
	require.Nil(t, err)

	// Assert that the returned habits match the expected values
	expectedHabits := []habit.Habit{
		{ID: "ID1", Name: "Knit", WeeklyFrequency: 3, Ticks: 3},
		{ID: "ID2", Name: "Code", WeeklyFrequency: 5, Ticks: 2},
	}

	assert.Equal(t, expectedHabits, habits)
}

func TestListHabits_error(t *testing.T) {
	sentinelErr := fmt.Errorf("haute leider nicht")

	// Create a mock for the API client
	mockClient := mocks.NewHabitsClientMock(t)
	mockClient.ListHabitsMock.Expect(minimock.AnyContext, &api.ListHabitsRequest{}).Return(nil, sentinelErr)

	habitsClient := client.New(mockClient)
	// Call the function being tested
	habits, err := habitsClient.ListHabits(context.Background(), time.Now())

	require.ErrorIs(t, err, sentinelErr)
	assert.Nil(t, habits)
}

func TestListHabits_statuserror(t *testing.T) {
	now := time.Now()

	// Create a mock for the API client
	mockClient := mocks.NewHabitsClientMock(t)

	// Create a HabitsClient with the mock client
	habitsClient := client.New(mockClient)

	// Define sample data for the mock response
	mockResponse := &api.ListHabitsResponse{
		Habits: []*api.Habit{
			{Id: "ID1", Name: "Knit", WeeklyFrequency: 3},
			{Id: "ID2", Name: "Code", WeeklyFrequency: 5},
		},
	}

	sentinelErr := status.Error(codes.Internal, "not after 10PM")
	mockClient.ListHabitsMock.Expect(minimock.AnyContext, &api.ListHabitsRequest{}).Return(mockResponse, nil)
	mockClient.GetHabitStatusMock.Set(
		func(ctx context.Context, in *api.GetHabitStatusRequest, opts ...grpc.CallOption) (*api.GetHabitStatusResponse, error) {
			switch {
			case in.HabitId == "ID1":
				return &api.GetHabitStatusResponse{}, nil
			case in.HabitId == "ID2":
				return nil, sentinelErr
			default:
				return nil, fmt.Errorf("unexpected call")
			}
		})

	// Call the function being tested
	habits, err := habitsClient.ListHabits(context.Background(), now)

	require.ErrorIs(t, err, sentinelErr)
	assert.Nil(t, habits)
}
