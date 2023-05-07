package handler

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AgentService struct {
	proto.UnimplementedAgentServiceServer
}

func (s *AgentService) ListAgents(ctx context.Context, in *emptypb.Empty) (*proto.AgentListResponse, error) {
	fmt.Println("[gRPC] [AgentService] ListAgents")
	return &proto.AgentListResponse{Success: true}, nil
}

func (s *AgentService) GetAgentInfo(ctx context.Context, in *proto.GetAgentInfoRequest) (*proto.GetAgentInfoResponse, error) {
	fmt.Println("[gRPC] [AgentService] GetAgentInfo:", in.Id)
	return &proto.GetAgentInfoResponse{Success: true}, nil
}

func (s *AgentService) ExecuteCmd(ctx context.Context, in *proto.ExecuteCmdRequest) (*proto.ExecuteCmdResponse, error) {
	fmt.Println("[gRPC] [AgentService] ExecuteCmd:", in.AgentId, "Cmd:", in.Cmd)
	return &proto.ExecuteCmdResponse{Success: true}, nil
}
