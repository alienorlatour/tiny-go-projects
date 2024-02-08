package client

import (
	"learngo-pockets/habits/api"
)

// HabitsClient is a wrapper around the gRPC client.
type HabitsClient struct {
	cli api.HabitsClient
}

func New(cli api.HabitsClient) *HabitsClient {
	return &HabitsClient{cli: cli}
}
