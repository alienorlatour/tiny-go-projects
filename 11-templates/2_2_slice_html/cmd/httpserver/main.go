package main

import (
	"fmt"
	"net/http"
	"os"

	"learngo-pockets/templates/internal/client"
	"learngo-pockets/templates/internal/handlers"
	"learngo-pockets/templates/internal/log"

	"learngo-pockets/habits/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = 8083

func main() {
	lgr := log.New(os.Stdout)

	cli, err := newClient("localhost:8084")
	if err != nil {
		lgr.Logf("Error while creating backend client: %s", err.Error())
	}

	srv := handlers.New(cli, lgr)

	addr := fmt.Sprintf(":%d", port)
	lgr.Logf("Listening on %d...", port)

	err = http.ListenAndServe(addr, srv.Router())
	if err != nil {
		panic(err)
	}
}

func newClient(serverAddress string) (*client.HabitsClient, error) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddress, creds)
	if err != nil {
		return nil, err
	}

	grpCli := api.NewHabitsClient(conn)

	return client.New(grpCli), nil
}
