package main

import (
	"authmaster/authmaster"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"authmaster/endpoints"
)

type server struct {
	authmaster.AuthmasterServer
}

const (
	AUTH_PREFIX = "Authorization"
)

func (s *server) TestAuth(ctx context.Context, in *authmaster.TestAuthRequest) (*authmaster.TestAuthResponse, error) {
	return endpoints.HandleTest(ctx)
}

func (s *server) CreateUser(ctx context.Context, in *authmaster.CreateUserRequest) (*authmaster.CreateUserResponse, error) {
	return endpoints.CreateUser(ctx, in)
}

func (s *server) Login(ctx context.Context, in *authmaster.LoginRequest) (*authmaster.LoginResponse, error) {
	return endpoints.Login(ctx, in)
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	authmaster.RegisterAuthmasterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
