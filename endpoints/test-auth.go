package endpoints

import (
	"context"
	"fmt"
	"time"

	"github.com/WadeCappa/authmaster/authmaster"
	"github.com/WadeCappa/authmaster/store"
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
		return nil, NO_AUTH_HEADER
	}

	t, ok := md["authorization"]
	if !ok {
		return nil, NO_AUTH_HEADER
	}

	if len(t) != 1 {
		return nil, NO_AUTH_HEADER
	}

	token := t[0]
	testResult, err := store.Call(func(conn *pgx.Conn) testResult {
		var userId int64
		var expireTime pgtype.Date
		err := conn.QueryRow(context.Background(), "select user_id, expire_time from tokens where token=$1", token).Scan(&userId, &expireTime)
		if err != nil {
			return testResult{e: fmt.Errorf("sql query failed: %v", err)}
		}
		return testResult{userId: userId, expireTime: expireTime, e: nil}
	})
	if err != nil {
		fmt.Printf("Failed to connect to postgres: %s\n", err)
		return nil, POSTGRES_CONNECTION_ERROR
	}
	if testResult.e != nil {
		fmt.Printf("Failed to return a result: %v\n", testResult.e)
		return nil, testResult.e
	}
	if testResult.expireTime.Time.Before(time.Now()) {
		fmt.Printf("Token has expired: %v\n", testResult)
		return nil, PERMISSION_DENIED
	}

	return &authmaster.TestAuthResponse{UserId: testResult.userId}, nil
}
