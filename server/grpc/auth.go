package grpc

import (
	"context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	fmt.Println("[gRPC] [AuthService] Login:", in.GetUsername())
	return &proto.LoginResponse{Success: true}, nil
}

func (s *AuthService) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println("[gRPC] [AuthService] Register:", in.GetUsername())
	return &proto.RegisterResponse{Success: true}, nil
}
