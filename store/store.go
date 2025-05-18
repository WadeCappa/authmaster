package store

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Call[T any](query func(*pgx.Conn) *T) (*T, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())
	return query(conn), nil
}
