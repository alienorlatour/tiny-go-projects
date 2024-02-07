package main

import (
	"context"
	"log"

	"learngo-pockets/habits/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:28710", creds)
	if err != nil {
		log.Fatal(err)
	}

	mock(api.NewHabitsClient(conn))
}

func mock(cli api.HabitsClient) {
	cli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            "Read",
		WeeklyFrequency: ptr(7),
	})
	cli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            "Feed the sourdough",
		WeeklyFrequency: ptr(14),
	})
}

func ptr(i int32) *int32 {
	return &i
}
