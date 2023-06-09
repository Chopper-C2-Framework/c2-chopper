package grpc

import (
	context "context"
	"errors"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/services"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AgentService struct {
	proto.UnimplementedAgentServiceServer

	AgentService services.IAgentService
}

func (s *AgentService) ListAgents(ctx context.Context, in *emptypb.Empty) (*proto.AgentListResponse, error) {
	fmt.Println("[gRPC] [AgentService] ListAgents")
	agents, err := s.AgentService.FindAllAgents()
	if err != nil {
		return &proto.AgentListResponse{}, err
	}

	protoList := make([]*proto.Agent, len(agents))
	for i, agent := range agents {
		protoList[i] = ConvertAgentToProto(agent)
	}

	return &proto.AgentListResponse{
		Data:  protoList,
		Count: uint32(len(protoList)),
	}, nil
}

func (s *AgentService) SetAgentNickname(ctx context.Context, in *proto.SetAgentNicknameRequest) (*proto.SetAgentNicknameResponse, error) {
	fmt.Println("[gRPC] [AgentService] SetAgentNickname:", in.GetAgentId())
	if len(in.GetAgentId()) == 0 {
		return &proto.SetAgentNicknameResponse{}, errors.New("agent id is required")
	}
	agent, err := s.AgentService.FindAgentOrError(in.GetAgentId())
	if err != nil {
		return &proto.SetAgentNicknameResponse{}, err
	}

	agent.Nickname = in.GetNickname()
	err = s.AgentService.SaveAgent(agent)
	if err != nil {
		return &proto.SetAgentNicknameResponse{}, err
	}

	return &proto.SetAgentNicknameResponse{}, nil
}

func (s *AgentService) GetAgentInfo(ctx context.Context, in *proto.GetAgentInfoRequest) (*proto.GetAgentInfoResponse, error) {
	fmt.Println("[gRPC] [AgentService] GetAgentInfo:", in.GetAgentId())
	if len(in.GetAgentId()) == 0 {
		return &proto.GetAgentInfoResponse{}, errors.New("Agent id is required")
	}
	agent, err := s.AgentService.FindAgentOrError(in.GetAgentId())
	if err != nil {
		return &proto.GetAgentInfoResponse{}, err
	}

	return &proto.GetAgentInfoResponse{
		Agent: ConvertAgentToProto(agent),
	}, nil
}

func (s *AgentService) Connect(ctx context.Context, in *proto.AgentConnectionRequest) (*proto.AgentConnectionResponse, error) {
	agentInfo := in.GetData()
	if agentInfo == nil {
		return &proto.AgentConnectionResponse{}, errors.New("Agent info missing")
	}

	agent, err := s.AgentService.ConnectAgent(
		agentInfo.GetId(),
		&entity.AgentModel{
			Username: agentInfo.GetUsername(),
			Uid:      agentInfo.GetUserId(),
			Hostname: agentInfo.GetHostname(),
			Cwd:      agentInfo.GetCwd(),
		},
	)
	if err != nil {
		return &proto.AgentConnectionResponse{}, err
	}
	return &proto.AgentConnectionResponse{Uuid: agent.ID.String()}, nil

}
