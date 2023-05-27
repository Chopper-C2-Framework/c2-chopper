package grpc

import (
	"context"
	"fmt"
	"github.com/chopper-c2-framework/c2-chopper/core/services"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
	AuthService services.IAuthService
}

func (a *AuthService) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	fmt.Println("[gRPC] [AuthService] Login:", in.GetUsername())

	token, err := a.AuthService.Login(in.GetUsername(), in.GetPassword())
	if err != nil {
		return &proto.LoginResponse{Success: false}, err
	}

	return &proto.LoginResponse{Success: true, Token: token}, nil
}

func (a *AuthService) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println("[gRPC] [AuthService] Register:", in.GetUsername())

	token, err := a.AuthService.Register(in.GetUsername(), in.GetPassword())

	if err != nil {
		return &proto.RegisterResponse{Success: false}, err
	}

	return &proto.RegisterResponse{Success: true, Token: token}, nil
}
