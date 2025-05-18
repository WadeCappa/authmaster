package endpoints

import (
	"authmaster/authmaster"
	"authmaster/store"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	NO_AUTH_HEADER         = status.Errorf(codes.InvalidArgument, "No authheader provided")
	FAILED_TO_GET_METADATA = status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	INVALID_AUTH           = status.Errorf(codes.PermissionDenied, "Permission Denied")
	INTERNAL_ERROR         = status.Errorf(codes.Internal, "Failed to connect to postgres")
)

type result struct {
	success bool
	e       error
}

func HandleTest(ctx context.Context) (*authmaster.TestAuthResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	if !ok {
		return nil, FAILED_TO_GET_METADATA
	}
	if t, ok := md["authorization"]; ok {
		if len(t) != 1 {
			return nil, NO_AUTH_HEADER
		}

		token := t[0]
		fmt.Printf("Testing auth token %s", token)
		result, err := store.Call(func(conn *pgx.Conn) *result {
			return &result{success: true, e: nil}
		})
		if err != nil {
			fmt.Printf("Failed to connect to postgres, %s", err)
			return nil, INTERNAL_ERROR
		}
		if !result.success && result.e == nil {
			return nil, INVALID_AUTH
		}
		if !result.success {
			fmt.Printf("Encountered some error, %s", result.e)
			return nil, INTERNAL_ERROR
		}
		return &authmaster.TestAuthResponse{UserId: 0}, nil
	}

	return nil, NO_AUTH_HEADER

}
