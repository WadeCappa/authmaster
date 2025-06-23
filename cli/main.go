package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"syscall"

	"github.com/WadeCappa/authman/cli/requests"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	NO_COMMAND      = "no-command"
	FAKE_USERNAME   = "test-user"
	FAKE_PASSWORD   = "test-password"
	TESTING_ADDRESS = "localhost:50051"

	CREATE_ACCOUNT_CMD = "create-account"
	LOGIN_CMD          = "login"
)

var commands = map[string]func(*grpc.ClientConn){
	"login":          login,
	"create-account": createAccount,
	"test":           test,
}

var (
	addr = flag.String("addr", TESTING_ADDRESS, "the address to connect to")
	cmd  = flag.String("cmd", NO_COMMAND, "choose one of the following; create-account, login, test")
)

type userCreds struct {
	username string
	password string
}

func main() {
	flag.Parse()

	command := commands[*cmd]
	if command == nil {
		log.Fatalf("Invalid command %s", *cmd)
	}

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true, // we do this because we self-sign internally. There's probably a better method for checking security
	})
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	command(conn)
}

func readUsernameAndPassword() (*userCreds, error) {
	var username string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}

	return &userCreds{username: username, password: string(password)}, nil
}

func login(conn *grpc.ClientConn) {
	creds, err := readUsernameAndPassword()
	if err != nil {
		log.Fatalf("Failed to read creadentials! %v\n", err)
	}

	token, err := requests.Login(conn, creds.username, creds.password)
	if err != nil {
		log.Fatalf("Failed to login! %v\n", err)
	}

	log.Printf("%s", *token)
}

func createAccount(conn *grpc.ClientConn) {
	creds, err := readUsernameAndPassword()
	if err != nil {
		log.Fatalf("Failed to read creadentials! %v\n", err)
	}

	if err := requests.CreateAccount(conn, creds.username, creds.password); err != nil {
		log.Fatalf("Failed to login! %v\n", err)
	}

	log.Printf("created new account for %s", creds.username)
}

func test(conn *grpc.ClientConn) {
	fmt.Print("Copy-paste your token:")
	token, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read token! %v\n", err)
	}

	userId, err := requests.Test(conn, string(token))
	if err != nil {
		log.Fatalf("Test failed %v", err)
	}

	fmt.Printf("test returned user id of %d\n", *userId)
}
