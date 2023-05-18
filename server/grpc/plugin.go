package grpc

import (
	context "context"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
)

type PluginService struct {
	proto.UnimplementedPluginServiceServer
	PluginManager plugins.IPluginManager
}

func (s *PluginService) ListLoadedPlugins(ctx context.Context, in *emptypb.Empty) (*proto.ListLoadedPluginsResponse, error) {
	fmt.Println("[gRPC] [PluginService] ListLoadedPlugins")
	allLoadedPlugins := s.PluginManager.ListLoadedPlugins()
	return &proto.ListLoadedPluginsResponse{Success: true, Names: allLoadedPlugins}, nil
}

func (s *PluginService) ListPlugins(ctx context.Context, in *emptypb.Empty) (*proto.ListPluginsResponse, error) {
	fmt.Println("[gRPC] [PluginService] ListPlugins")
	allPlugins, err := s.PluginManager.ListAllPlugins()
	if err != nil {
		return &proto.ListPluginsResponse{Success: false}, err
	}
	return &proto.ListPluginsResponse{Success: true, Names: allPlugins}, nil
}

func (s *PluginService) RunPlugin(ctx context.Context, in *proto.RunPluginRequest) (*proto.RunPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] RunPlugin")
	plugin, err := s.PluginManager.GetPlugin(in.FileName)
	if err != nil {
		return &proto.RunPluginResponse{Success: false}, err
	}

	var args []interface{}
	for _, v := range in.Args {
		args = append(args, GetValue(v))
	}

	// Run plugin & return result
	err = plugin.SetArgs(args...)
	if err != nil {
		return &proto.RunPluginResponse{Success: false}, err
	}
	result := plugin.Exploit()
	return &proto.RunPluginResponse{Success: true, Message: string(result)}, nil
}

func (s *PluginService) LoadPlugin(ctx context.Context, in *proto.LoadPluginRequest) (*proto.LoadPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] LoadPlugin")
	plugin, err := s.PluginManager.LoadPlugin(in.FileName)
	if err != nil {
		return &proto.LoadPluginResponse{Success: false}, err
	}
	return &proto.LoadPluginResponse{Success: true, Data: s.GetPluginInfo(plugin)}, nil
}

func (s *PluginService) GetPluginDetails(ctx context.Context, in *proto.LoadPluginRequest) (*proto.LoadPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] GetPluginDetails")
	plugin, err := s.PluginManager.GetPlugin(in.FileName)
	if err != nil {
		return &proto.LoadPluginResponse{Success: false}, err
	}
	return &proto.LoadPluginResponse{Success: true, Data: s.GetPluginInfo(plugin)}, nil
}

func (s *PluginService) GetPluginInfo(plugin plugins.IPlugin) *proto.Plugin {
	pluginInfo := plugin.Info()
	pluginMeta := plugin.MetaInfo()

	info := &proto.PluginInfo{
		Name:       pluginInfo.Name,
		Options:    pluginInfo.Options,
		ReturnType: pluginInfo.ReturnType,
	}
	metadata := &proto.PluginMetadata{
		Version:     pluginMeta.Version,
		Author:      pluginMeta.Author,
		Tags:        pluginMeta.Tags,
		ReleaseDate: pluginMeta.ReleaseDate,
		Type:        int32(pluginMeta.Type),
		SourceLink:  pluginMeta.SourceLink,
		Description: pluginMeta.Description,
	}
	return &proto.Plugin{Info: info, Metadata: metadata}
}

func GetValue(val *proto.ArgValue) interface{} {
	if val.Type == "string_value" {
		return val.GetStringValue()
	}
	if val.Type == "bool_value" {
		return val.GetBoolValue()
	}
	if val.Type == "number_value" {
		return val.GetNumberValue()
	}
	if val.Type == "map_value" {
		/*
			{
				"arg0": {
					"type": "string_value",
					"string_value": "value 1"
				},
				"arg1": {
					"type": "map_value",
					"map_value": {
						"items": [
							{
								"key": "key1",
								"value": {
									"type": "string_value",
									"string_value": "Hello world"
								}
							}
						]
					}
				}
			}
		*/
		var mapVariable map[string]interface{} = make(map[string]interface{})
		mapItems := val.GetMapValue().GetItems()
		for _, item := range mapItems {
			mapVariable[item.GetKey()] = GetValue(item.GetValue())
		}
		return mapVariable
	}
	return nil
}
