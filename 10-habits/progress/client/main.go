package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"learngo-pockets/habits/api"
)

func main() {
	ctx := context.Background()

	serverAddress := ":38804"
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error(fmt.Sprintf("error while creating the client: %s", err.Error()))
		os.Exit(1)
	}

	habitsCli := api.NewHabitsClient(conn)

	walkFrequency := int32(5)

	resp, err := habitsCli.CreateHabit(ctx, &api.CreateHabitRequest{Habit: &api.Habit{Name: "walk in forest", Frequency: &walkFrequency}})
	if err != nil {
		slog.Error(fmt.Sprintf("unexpected error while creating a habit: %s", err.Error()))
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("Response received: %s", resp))
}
