package handler

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ListenerService struct {
	proto.UnimplementedListenerServiceServer
}

func (s *ListenerService) ListListeners(ctx context.Context, in *emptypb.Empty) (*proto.ListenerListResponse, error) {
	fmt.Println("[gRPC] [ListenerService] ListListeners")
	return &proto.ListenerListResponse{Success: true}, nil
}

func (s *ListenerService) GetListenerInfo(ctx context.Context, in *proto.GetListenerInfoRequest) (*proto.GetListenerInfoResponse, error) {
	fmt.Println("[gRPC] [ListenerService] GetListenerInfo:", in.Id)
	return &proto.GetListenerInfoResponse{Success: true}, nil
}

func (s *ListenerService) ExecuteCmd(ctx context.Context, in *proto.ExecuteCmdRequest) (*proto.ExecuteCmdResponse, error) {
	fmt.Println("[gRPC] [ListenerService] ExecuteCmd:", in.ListenerId, "Cmd:", in.Cmd)
	return &proto.ExecuteCmdResponse{Success: true}, nil
}
