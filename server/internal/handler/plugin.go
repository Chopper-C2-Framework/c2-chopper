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
	plugins := s.PluginManager.ListLoadedPlugins()
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
	plugin, err := s.PluginManager.GetPlugin(in.FileName)
	if err != nil {
		return &proto.RunPluginResponse{Success: false}, err
	}
	// Run plugin & return result
	fmt.Println("[gRPC] [PluginService] RunPlugin ", plugin.Name)
	return &proto.RunPluginResponse{Success: true}, nil
}

func (s *PluginService) LoadPlugin(ctx context.Context, in *proto.LoadPluginRequest) (*proto.LoadPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] LoadPlugin")
	plugin, err := s.PluginManager.LoadPlugin(in.FileName)
	if err != nil {
		return &proto.LoadPluginResponse{Success: false}, err
	}
	return &proto.LoadPluginResponse{Success: true, Data: s.GetPluginInfo(plugin)}, nil
}

func (s *PluginService) GetPluginInfo(plugin *plugins.Plugin) *proto.Plugin {
	info := &proto.PluginInfo{
		Options:    plugin.PluginInfo.Options,
		ReturnType: plugin.PluginInfo.ReturnType,
	}
	metadata := &proto.PluginMetadata{
		Version:     plugin.Version,
		Author:      plugin.Author,
		Tags:        plugin.Tags,
		ReleaseDate: plugin.ReleaseDate,
		Type:        int32(plugin.Type),
		SourceLink:  plugin.SourceLink,
		Description: plugin.Description,
	}
	return &proto.Plugin{Name: plugin.Name, Info: info, Metadata: metadata}
}
