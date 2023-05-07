package handler

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PluginService struct {
	proto.UnimplementedPluginServiceServer
}

func (s *PluginService) ListLoadedPlugins(ctx context.Context, in *emptypb.Empty) (*proto.ListPluginsResponse, error) {
	fmt.Println("[gRPC] [PluginService] ListLoadedPlugins")
	return &proto.ListPluginsResponse{Success: true}, nil
}

func (s *PluginService) ListPlugins(ctx context.Context, in *emptypb.Empty) (*proto.ListPluginsResponse, error) {
	fmt.Println("[gRPC] [PluginService] ListPlugins")
	return &proto.ListPluginsResponse{Success: true}, nil
}

func (s *PluginService) RunPlugin(ctx context.Context, in *proto.RunPluginRequest) (*proto.RunPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] RunPlugin")
	return &proto.RunPluginResponse{Success: true}, nil
}

func (s *PluginService) LoadPlugin(ctx context.Context, in *proto.LoadPluginRequest) (*proto.LoadPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] LoadPlugin")
	return &proto.LoadPluginResponse{Success: true}, nil
}
