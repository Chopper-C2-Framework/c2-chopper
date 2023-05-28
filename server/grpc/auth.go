package grpc

import (
	"context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/services"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"google.golang.org/grpc/metadata"
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

func (a *AuthService) Me(ctx context.Context, in *proto.MeRequest) (*proto.MeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &proto.MeResponse{
			Success: false,
			User:    nil,
			Teams:   nil,
		}, nil
	}

	fmt.Println(md)
	userIds := md.Get("userid")

	fmt.Println("[gRPC] [AuthService] Me:", userIds)

	if len(userIds) > 0 {

		user, err := a.AuthService.FetchUserFromId(userIds[0])
		if err != nil {
			return &proto.MeResponse{
				Success: false,
			}, nil
		}

		fmt.Println("Found userId", userIds[0])
		return &proto.MeResponse{
			Success: true,
			User:    ConvertUserToProto(user),
			Teams:   ConvertTeamsToProto(user.Teams),
		}, nil

	}

	return &proto.MeResponse{
		Success: false,
	}, nil

}
