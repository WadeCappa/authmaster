package endpoints

import (
	"authmaster/authmaster"
	"authmaster/store"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/metadata"
)

type testResult struct {
	userId     int64
	expireTime pgtype.Date
	e          error
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
		testResult, err := store.Call(func(conn *pgx.Conn) *testResult {
			var userId int64
			var expireTime pgtype.Date
			err := conn.QueryRow(context.Background(), "select user_id, expire_time from users where token=$1", token).Scan(&userId, &expireTime)
			if err != nil {
				fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
				os.Exit(1)
			}
			return &testResult{userId: userId, expireTime: expireTime, e: nil}
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to connect to postgres, %s", err)
			return nil, INTERNAL_ERROR
		}
		if testResult.e != nil {
			fmt.Fprintf(os.Stderr, "Encountered some error, %s", testResult.e)
			return nil, INTERNAL_ERROR
		}
		if testResult.expireTime.Time.Before(time.Now()) {
			return nil, INVALID_AUTH
		}
		return &authmaster.TestAuthResponse{UserId: int32(testResult.userId)}, nil
	}

	return nil, NO_AUTH_HEADER

}
