package main

import (
	"fmt"
	"log"
	"net/http"

	"learngo-pockets/templates/internal/handlers"
)

const port = 8083

func main() {
	cli, err := newClient("localhost:28710")
	if err != nil {
		log.Fatalf("Error while creating backend client: %s", err.Error())
	}

	srv := handlers.New(cli)

	addr := fmt.Sprintf(":%d", port)
	log.Print("Listening on ", port, "...")

	err = http.ListenAndServe(addr, srv.Router())
	if err != nil {
		panic(err)
	}
}

// newClient creates a client that can call the grpc habits service.
func newClient(serverAddress string) (*client.HabitsClient, error) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddress, creds)
	if err != nil {
		return nil, err
	}

	grpCli := api.NewHabitsClient(conn)

	return client.New(grpCli), nil
}
