package server

import (
	"github.com/WadeCappa/authmaster/pkg/go/authmaster/v1"
)

type server struct {
	authmaster.AuthmasterServer

	postgresUrl string
}

func NewServer(postgresUrl string) server {
	return server{postgresUrl: postgresUrl}
}
