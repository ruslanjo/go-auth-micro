package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/ruslanjo/go-auth-micro/pkg/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	desc.UnimplementedUsersV1Server
}

func (s *Server) Create(ctx context.Context, data *desc.CreateRequest) (*desc.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id %d", req.GetId())
	return &desc.GetResponse{
		User: &desc.User{
			Id:   req.GetId(),
			Info: &desc.PublicUserInfo{
				Login: "ruslanjo",
				Country: "Russia",
			},
			CreatedAt: timestamppb.New(
				time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
			),
		},
	}, nil
}
func (s *Server) List(context.Context, *desc.ListRequest) (*desc.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (s *Server) Update(context.Context, *desc.UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *Server) Delete(context.Context, *desc.DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
