package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/WadeCappa/authmaster/authmaster"
	"github.com/WadeCappa/authmaster/server"
	"google.golang.org/grpc"
)

const (
	AUTH_PREFIX                  = "Authorization"
	TESTING_DEFAULT_POSTGRES_URL = "postgres://postgres:pass@postgres:5432/authmaster_db"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	postgresUrl = flag.String("postgresUrl", TESTING_DEFAULT_POSTGRES_URL, "the connection URL to our postgres instance")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := server.NewServer(*postgresUrl)
	authmaster.RegisterAuthmasterServer(s, &server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
