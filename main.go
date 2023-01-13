package main

import (
	"log"
	"net"

	"github.com/barretodotcom/go-grpc-integration/infra/db"
	"github.com/barretodotcom/go-grpc-integration/infra/provider"
	"github.com/barretodotcom/go-grpc-integration/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Couldn't get .env file")
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3333")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServer(server, &provider.UserServer{})
	pb.RegisterAuthServer(server, &provider.AuthServer{})

	server.Serve(listener)

}
