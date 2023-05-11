package grpc

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
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

func (s *TeamService) AddMemberToTeam(ctx context.Context, in *proto.AddMemberRequest) (*proto.AddMemberResponse, error) {

	fmt.Println("[gRPC] [TeamService] AddMemberToTeam:", in.GetTeamId())

	return nil, nil
}

func (s *TeamService) UpdateTeam(ctx context.Context, in *proto.UpdateTeamRequest) (*proto.UpdateTeamResponse, error) {

	fmt.Println("[gRPC] [TeamService] UpdateTeam:", in.GetTeamId())

	return nil, nil
}

func (s *TeamService) DeleteTeam(ctx context.Context, in *proto.DeleteTeamRequest) (*proto.DeleteTeamResponse, error) {

	fmt.Println("[gRPC] [TeamService] DeleteTeam:", in.GetTeamId())

	return nil, nil
}
