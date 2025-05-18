package main

import (
	"context"
	"flag"
	"log"

	"authmaster/authmaster"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
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

	r, err := c.TestAuth(ctx, &authmaster.TestAuthRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %d", r.UserId)
}
