package requests

import (
	"context"

	"github.com/WadeCappa/authmaster/authmaster"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Login(conn *grpc.ClientConn, username, password string) (*string, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs())
	c := authmaster.NewAuthmasterClient(conn)

	response, err := c.Login(ctx, &authmaster.LoginRequest{Username: username, Password: string(password)})
	if err != nil {
		return nil, err
	}

	return &response.Token, nil
}
