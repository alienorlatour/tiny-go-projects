package client

import (
	"context"
	"fmt"
	"testing"

	"learngo-pockets/habits/api"
	"learngo-pockets/templates/internal/client/mocks"
	habit "learngo-pockets/templates/internal/habits"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

//go:generate minimock -i learngo-pockets/habits/api.HabitsClient -s "_mock.go" -o "mocks"

func TestListHabits(t *testing.T) {
	// Create a mock for the API client
	mockClient := mocks.NewHabitsClientMock(t)

	// Create a HabitsClient with the mock client
	habitsClient := New(mockClient)

	// Define sample data for the mock response
	mockResponse := &api.ListHabitsResponse{
		Habits: []*api.Habit{
			{Id: "1", Name: "Reading", WeeklyFrequency: 3},
			{Id: "2", Name: "Exercise", WeeklyFrequency: 5},
		},
	}

	mockClient.ListHabitsMock.Expect(minimock.AnyContext, &api.ListHabitsRequest{}).Return(mockResponse, nil)

	// Call the function being tested
	habits, err := habitsClient.ListHabits(context.Background())

	// Assert that there are no errors
	assert.Nil(t, err)

	// Assert that the returned habits match the expected values
	expectedHabits := []habit.Habit{
		{ID: "1", Name: "Reading", WeeklyFrequency: 3},
		{ID: "2", Name: "Exercise", WeeklyFrequency: 5},
	}

	assert.Equal(t, expectedHabits, habits)
}

func TestListHabits_error(t *testing.T) {
	sentinelErr := fmt.Errorf("haute leider nicht")

	// Create a mock for the API client
	mockClient := mocks.NewHabitsClientMock(t)
	mockClient.ListHabitsMock.Expect(minimock.AnyContext, &api.ListHabitsRequest{}).Return(nil, sentinelErr)

	habitsClient := New(mockClient)
	// Call the function being tested
	habits, err := habitsClient.ListHabits(context.Background())

	assert.ErrorIs(t, err, sentinelErr)
	assert.IsNonDecreasing(t, habits)
}
