package endpoints

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"encoding/base64"

	"github.com/WadeCappa/authmaster/authmaster"
	"github.com/WadeCappa/authmaster/hashing"
	"github.com/WadeCappa/authmaster/store"
	"github.com/jackc/pgx/v5"
)

const (
	keepAliveHours  = 72
	authTokenLength = 256
)

type getUserDataResult struct {
	err    error
	userId *int64
	dbHash *string
}

type makeNewTokenResult struct {
	err      error
	newToken *string
}

type getPreviousTokenResult struct {
	err           error
	previousToken *string
}

// currently we have a race on checking the password and making a new token.
// This is fine for now since user data is completely immutable, but in the
// future if we want to support users changing their passwords, we'll have to
// fix this.
func Login(ctx context.Context, in *authmaster.LoginRequest) (*authmaster.LoginResponse, error) {

	userData, err := store.Call(func(conn *pgx.Conn) getUserDataResult {
		var dbHash string
		var userId int64
		err := conn.QueryRow(context.Background(), "select user_id, hash from users where username=$1", in.Username).Scan(&userId, &dbHash)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to find password hash, %s", err)
			return getUserDataResult{err: err, userId: nil}
		}
		return getUserDataResult{err: nil, userId: &userId, dbHash: &dbHash}
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to postgres, %s\n", err)
		return nil, INTERNAL_ERROR
	}

	if userData.userId == nil || userData.dbHash == nil {
		return nil, PERMISSION_DENIED
	}

	if hashing.VerifyPassword(in.Password, *userData.dbHash) != nil {
		fmt.Fprintf(os.Stderr, "Failed to verify password, %s\n", err)
		return nil, PERMISSION_DENIED
	}

	previousToken, err := checkForPreviousToken(*userData.userId)
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	if previousToken != nil {
		return &authmaster.LoginResponse{Token: *previousToken}, nil
	}

	token, err := makeToken(authTokenLength)
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	expireTime := time.Now().Add(time.Duration(keepAliveHours * time.Hour))

	newTokenResult, err := store.Call(func(conn *pgx.Conn) makeNewTokenResult {
		var newToken string
		err := conn.QueryRow(context.Background(), "insert into tokens (user_id, token, expire_time) values ($1, $2, $3) returning token", userData.userId, token, expireTime).Scan(&newToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to store new token %s\n", err)
			return makeNewTokenResult{err: err, newToken: nil}
		}
		return makeNewTokenResult{err: nil, newToken: &newToken}
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create new token, %s\n", err)
		return nil, INTERNAL_ERROR
	}

	if newTokenResult.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create new token, %s\n", newTokenResult.err)
		return nil, INTERNAL_ERROR
	}

	return &authmaster.LoginResponse{Token: *newTokenResult.newToken}, nil
}

func makeToken(n int) (*string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create token, %s\n", err)
		return nil, err
	}

	encoded := base64.RawURLEncoding.EncodeToString(bytes)
	shortened := encoded[0:authTokenLength]
	fmt.Printf("created token, not yet stored. Has length of %d\n", len(shortened))
	return &shortened, nil
}

func checkForPreviousToken(userId int64) (*string, error) {
	result, err := store.Call(func(conn *pgx.Conn) getPreviousTokenResult {
		var token string
		err := conn.QueryRow(context.Background(), "select token from tokens where user_id=$1", userId).Scan(&token)
		if err != nil {
			return getPreviousTokenResult{err: err, previousToken: nil}
		}
		return getPreviousTokenResult{previousToken: &token, err: nil}
	})

	if result.err != nil {
		return nil, err
	}

	return result.previousToken, nil
}
