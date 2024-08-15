package main

import (
	"context"
	"log"
	"time"

	desc "github.com/ruslanjo/go-auth-micro/pkg/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcAddr string = "localhost:60001"
	userID   int64  = 3
)

func main() {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server %v", err)
	}

	defer conn.Close()

	client := desc.NewUsersV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := client.Get(ctx, &desc.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("failed to get user by id %v", err)
	}

	log.Printf("%+v", req.GetUser())

}
