package handler

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
)

type PluginService struct {
	proto.UnimplementedPluginServiceServer
	PluginManager plugins.IPluginManager
}

func (s *PluginService) ListLoadedPlugins(ctx context.Context, in *emptypb.Empty) (*proto.ListPluginsResponse, error) {
	fmt.Println("[gRPC] [PluginService] ListLoadedPlugins")
	plugins, _ := s.PluginManager.ListLoadedPlugins()
	return &proto.ListPluginsResponse{Success: true, Names: plugins}, nil
}

func (s *PluginService) ListPlugins(ctx context.Context, in *emptypb.Empty) (*proto.ListPluginsResponse, error) {
	fmt.Println("[gRPC] [PluginService] ListPlugins")
	plugins, err := s.PluginManager.ListAllPlugins()
	if err != nil {
		return &proto.ListPluginsResponse{Success: false}, err
	}
	return &proto.ListPluginsResponse{Success: true, Names: plugins}, nil
}

func (s *PluginService) RunPlugin(ctx context.Context, in *proto.RunPluginRequest) (*proto.RunPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] RunPlugin")
	return &proto.RunPluginResponse{Success: true}, nil
}

func (s *PluginService) LoadPlugin(ctx context.Context, in *proto.LoadPluginRequest) (*proto.LoadPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] LoadPlugin")
	fmt.Println(in)
	plugin, err := s.PluginManager.LoadPlugin(in.FileName)
	if err != nil {
		return &proto.LoadPluginResponse{Success: false}, err
	}
	return &proto.LoadPluginResponse{Success: true, Data: &proto.Plugin{Name: plugin.Name}}, nil
}
