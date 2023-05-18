package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/services"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

type TeamService struct {
	proto.UnimplementedTeamServiceServer
	TeamService services.ITeamService
}

func (s *TeamService) CreateTeam(ctx context.Context, in *proto.CreateTeamRequest) (*proto.CreateTeamResponse, error) {
	fmt.Println("[gRPC] [TeamService] CreateTeam:", in.Data.GetId())

	incomingData := in.GetData()
	if incomingData == nil {
		return &proto.CreateTeamResponse{Success: false}, errors.New("unable to create team, no data provided")
	}

	newTeam := &entity.TeamModel{
		Name:    in.GetData().Name,
		Members: []*entity.UserModel{},
	}

	err := s.TeamService.CreateTeam(newTeam)
	if err != nil {
		return &proto.CreateTeamResponse{Success: false}, err
	}

	return &proto.CreateTeamResponse{Success: true, Data: ConvertTeamToProto(newTeam)}, nil
}

func (s *TeamService) JoinTeam(ctx context.Context, in *proto.JoinTeamRequest) (*proto.JoinTeamResponse, error) {

	fmt.Println("[gRPC] [TeamService] JoinTeam:", in.GetTeamId())

	teamId := in.GetTeamId()

	if teamId == "" {
		return &proto.JoinTeamResponse{Success: false}, errors.New("unable to create team, no data provided")
	}

	// Get user id from context
	userId := ctx.Value("userId")
	if userId == nil {
		return &proto.JoinTeamResponse{Success: false}, nil
	}
	err := s.TeamService.AddMemberToTeam(in.GetTeamId(), userId.(string))

	if err != nil {
		return &proto.JoinTeamResponse{Success: false}, err
	}

	return &proto.JoinTeamResponse{Success: true}, nil
}

func (s *TeamService) AddMemberToTeam(ctx context.Context, in *proto.AddMemberToTeamRequest) (*proto.AddMemberToTeamResponse, error) {

	// VerifyAdmin()
	fmt.Println("[gRPC] [TeamService] AddMemberToTeam:", in.GetTeamId())

	err := s.TeamService.AddMemberToTeam(in.GetTeamId(), in.GetUserId())
	if err != nil {
		return &proto.AddMemberToTeamResponse{Success: false}, err
	}

	team, err := s.TeamService.FindOne(in.GetTeamId())

	return &proto.AddMemberToTeamResponse{Success: true, Team: ConvertTeamToProto(team)}, nil
}

func (s *TeamService) UpdateTeam(ctx context.Context, in *proto.UpdateTeamRequest) (*proto.UpdateTeamResponse, error) {

	fmt.Println("[gRPC] [TeamService] UpdateTeam:", in.GetTeamId())

	teamId := in.GetTeamId()

	updatedTeam := &entity.TeamModel{
		Name: in.UpdatedTeam.Name,
	}

	updatedTeam, err := s.TeamService.UpdateTeam(teamId, updatedTeam)

	if err != nil {
		return &proto.UpdateTeamResponse{Success: false}, err
	}

	return &proto.UpdateTeamResponse{Success: true, Data: ConvertTeamToProto(updatedTeam)}, nil

}

func (s *TeamService) DeleteTeam(ctx context.Context, in *proto.DeleteTeamRequest) (*proto.DeleteTeamResponse, error) {

	err := s.TeamService.DeleteTeam(in.GetTeamId())

	if err != nil {
		return &proto.DeleteTeamResponse{Success: false}, err
	}

	fmt.Println("[gRPC] [TeamService] DeleteTeam:", in.GetTeamId())

	return &proto.DeleteTeamResponse{Success: true}, nil
}
