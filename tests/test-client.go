package main

import (
	"authmaster/authmaster"
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

const (
	fakeUsername = "test-user"
	fakePassword = "test-password"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := authmaster.NewAuthmasterClient(conn)

	// Create metadata and context.
	md := metadata.Pairs("Authorization", "fake-token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	response, expectedError := c.TestAuth(ctx, &authmaster.TestAuthRequest{})
	log.Printf("failed to test fake auth token (expected): %v\n", expectedError)
	log.Printf("had response of %s", response.String())

	create, err := c.CreateUser(ctx, &authmaster.CreateUserRequest{Username: fakeUsername, Password: fakePassword})
	if err != nil {
		log.Fatalf("could not create user: %v\n", err)
	}

	log.Printf("created user %s\n", create.String())

	login, err := c.Login(ctx, &authmaster.LoginRequest{Username: fakeUsername, Password: fakePassword})
	if err != nil {
		log.Fatalf("could not login: %v\n", err)
	}

	log.Printf("logged in with token %s", login.Token)

	newMetadata := metadata.Pairs("Authorization", login.Token)
	newContext := metadata.NewOutgoingContext(context.Background(), newMetadata)

	shouldPass, err := c.TestAuth(newContext, &authmaster.TestAuthRequest{})

	if err != nil {
		log.Fatalf("Failed to test auth: %v\n", err)
	}
	log.Printf("response of %s", shouldPass.String())
}
