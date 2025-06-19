package server

import (
	"github.com/WadeCappa/authmaster/authmaster"
)

type server struct {
	authmaster.AuthmasterServer

	postgresUrl string
}

func NewServer(postgresUrl string) server {
	return server{postgresUrl: postgresUrl}
}
