package handler

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/proto"
)

type TeamService struct {
	proto.UnimplementedTeamServiceServer
}

func (s *TeamService) CreateTeam(ctx context.Context, in *proto.CreateTeamRequest) (*proto.CreateTeamResponse, error) {
	fmt.Println("[gRPC] [TeamService] CreateTeam:", in.Data.GetId())
	return &proto.CreateTeamResponse{Success: true}, nil
}

func (s *TeamService) JoinTeam(ctx context.Context, in *proto.JoinTeamRequest) (*proto.JoinTeamResponse, error) {
	fmt.Println("[gRPC] [TeamService] JoinTeam:", in.GetTeamId())
	return &proto.JoinTeamResponse{Success: true}, nil
}
