package main

import (
	"authmaster/authmaster"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	authmaster.AuthmasterServer
}

const AUTH_PREFIX = "Authorization"

func (s *server) TestAuth(ctx context.Context, in *authmaster.TestAuthRequest) (*authmaster.TestAuthResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["authorization"]; ok {
		fmt.Printf("Testing auth token %s", t)
	}
	return &authmaster.TestAuthResponse{UserId: 0}, nil
}

func (s *server) CreateUser(_ context.Context, in *authmaster.CreateUserRequest) (*authmaster.CreateUserResponse, error) {
	return &authmaster.CreateUserResponse{}, nil
}

func (s *server) Login(_ context.Context, in *authmaster.LoginRequest) (*authmaster.LoginResponse, error) {
	return &authmaster.LoginResponse{Token: "fake-token"}, nil
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

// func getToken(c *gin.Context) string {
// 	return c.Request.Header[AUTH_PREFIX]
// }
