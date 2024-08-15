package main

import (
	"fmt"
	"log"
	"net"

	desc "github.com/ruslanjo/go-auth-micro/pkg/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort int = 60_001

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on %d: %w", grpcPort, err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUsersV1Server(s, &Server{})

	log.Printf("server listening at %d", listen.Addr())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve %w", err)
	}
}
