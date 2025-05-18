package endpoints

import (
	"authmaster/authmaster"
	"authmaster/store"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type createUserResult struct {
	err error
}

func CreateUser(ctx context.Context, in *authmaster.CreateUserRequest) (*authmaster.CreateUserResponse, error) {

	result, err := store.Call(func(conn *pgx.Conn) *createUserResult {
		var newUserId int64
		conn.QueryRow(context.Background(), "select nextval('user_ids')").Scan(&newUserId)
		tag, err := conn.Exec(context.Background(), "insert into users (username, hash, user_id) values ($1, $2, $3)", in.Username, hashPassword(in.Password), newUserId)
		fmt.Println("tag from new user request, %s", tag)
		if err != nil {
			return &createUserResult{err: err}
		}
		return &createUserResult{err: nil}
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to postgres, %s", err)
		return nil, INTERNAL_ERROR
	}
	if result.err != nil {
		fmt.Fprintf(os.Stderr, "Encountered an error while creating a new account", result.err)
		return nil, INTERNAL_ERROR
	}
	return &authmaster.CreateUserResponse{}, nil
}

func hashPassword(password string) string {

}
