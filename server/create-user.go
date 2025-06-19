package server

import (
	"context"
	"fmt"
	"os"

	"github.com/WadeCappa/authmaster/authmaster"
	"github.com/WadeCappa/authmaster/hashing"
	"github.com/WadeCappa/authmaster/store"
	"github.com/jackc/pgx/v5"
)

type createUserResult struct {
	err error
}

func (s *server) CreateUser(ctx context.Context, in *authmaster.CreateUserRequest) (*authmaster.CreateUserResponse, error) {
	newHash, err := hashing.HashPassword(in.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to hash password, %s\n", err)
		return nil, INTERNAL_ERROR
	}

	result, err := store.Call(s.postgresUrl, func(conn *pgx.Conn) createUserResult {
		var newUserId int64
		err := conn.QueryRow(context.Background(), "select nextval('user_ids')").Scan(&newUserId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get a new user id, %s\n", err)
			return createUserResult{err: err}
		}
		tag, err := conn.Exec(context.Background(), "insert into users (username, hash, user_id) values ($1, $2, $3)", in.Username, newHash, newUserId)
		fmt.Printf("tag from new user request, %s\n", tag)
		if err != nil {
			return createUserResult{err: err}
		}
		return createUserResult{err: nil}
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to postgres, %s\n", err)
		return nil, INTERNAL_ERROR
	}
	if result.err != nil {
		fmt.Fprintf(os.Stderr, "Encountered an error while creating a new account, %s\n", result.err)
		return nil, INTERNAL_ERROR
	}
	return &authmaster.CreateUserResponse{}, nil
}
