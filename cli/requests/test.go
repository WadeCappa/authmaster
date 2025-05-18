package requests

import (
	"context"

	"github.com/WadeCappa/authmaster/authmaster"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Test(conn *grpc.ClientConn, token string) (*int32, error) {

	newMetadata := metadata.Pairs("Authorization", token)
	newContext := metadata.NewOutgoingContext(context.Background(), newMetadata)

	c := authmaster.NewAuthmasterClient(conn)
	response, err := c.TestAuth(newContext, &authmaster.TestAuthRequest{})

	if err != nil {
		return nil, err
	}

	return &response.UserId, nil
}
