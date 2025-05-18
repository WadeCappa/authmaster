package requests

import (
	"context"

	"github.com/WadeCappa/authmaster/authmaster"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func CreateAccount(conn *grpc.ClientConn, username, password string) error {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs())
	c := authmaster.NewAuthmasterClient(conn)

	_, err := c.CreateUser(ctx, &authmaster.CreateUserRequest{Username: username, Password: string(password)})
	if err != nil {
		return err
	}

	return nil
}
