package grpc

import (
	context "context"
	"errors"
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
	"github.com/chopper-c2-framework/c2-chopper/core/services"
)

type PluginService struct {
	proto.UnimplementedPluginServiceServer
	PluginManager       plugins.IPluginManager
	PluginResultService services.IPluginResultService
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
	loadedPlugin, err := s.PluginManager.GetPlugin(in.GetFileName())
	if err != nil {
		return &proto.RunPluginResponse{Success: false}, err
	}

	if loadedPlugin.Channel != nil {
		return &proto.RunPluginResponse{}, errors.New("plugin is already executing.")
	}
	plugin := loadedPlugin.Plugin

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("userid")) == 0 {
		return &proto.RunPluginResponse{}, errors.New("unable to process")
	}
	userUUID, err := uuid.Parse(md.Get("userid")[0])
	if err != nil {
		return &proto.RunPluginResponse{Success: false}, err
	}

	args := make(map[string]interface{})
	for key, v := range in.Args {
		fmt.Println(key)
		args[key] = GetValue(v)
	}

	// Set plugin arguments
	err = plugin.SetArgs(args)
	if err != nil {
		return &proto.RunPluginResponse{Success: false}, err
	}

	// Start execution in a new thread. Store in db when result is ready.
	loadedPlugin.Channel = make(chan *entity.TaskResultModel, 1)
	go func() {
		result := plugin.Exploit(loadedPlugin.Channel)
		s.PluginResultService.CreatePluginResult(&entity.PluginResultModel{
			Output:     string(result),
			Path:       in.GetFileName(),
			OutputType: "string",
			CreatorId:  userUUID,
		})
		loadedPlugin.Channel = nil
	}()
	return &proto.RunPluginResponse{Success: true, Message: "WIP"}, nil
}

func (s *PluginService) LoadPlugin(ctx context.Context, in *proto.LoadPluginRequest) (*proto.LoadPluginResponse, error) {
	fmt.Println("[gRPC] [PluginService] LoadPlugin")
	loadedPlugin, err := s.PluginManager.LoadPlugin(in.FileName)
	if err != nil {
		return &proto.LoadPluginResponse{Success: false}, err
	}
	plugin := loadedPlugin.Plugin

	return &proto.LoadPluginResponse{Success: true, Data: ConvertPluginToProto(plugin)}, nil
}

func (s *PluginService) GetPluginDetails(ctx context.Context, in *proto.GetPluginDetailsRequest) (*proto.GetPluginDetailsResponse, error) {
	fmt.Println("[gRPC] [PluginService] GetPluginDetails")
	loadedPlugin, err := s.PluginManager.GetPlugin(in.FileName)
	if err != nil {
		return &proto.GetPluginDetailsResponse{}, err
	}
	plugin := loadedPlugin.Plugin

	return &proto.GetPluginDetailsResponse{Data: ConvertPluginToProto(plugin)}, nil
}

func (s *PluginService) GetPluginResults(ctx context.Context, in *proto.GetPluginResultsRequest) (*proto.GetPluginResultsResponse, error) {
	fmt.Println("[gRPC] [PluginService] GetPluginResults")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("userid")) == 0 {
		return &proto.GetPluginResultsResponse{}, errors.New("unable to process")
	}
	userUUID, err := uuid.Parse(md.Get("userid")[0])
	if err != nil {
		return &proto.GetPluginResultsResponse{}, err
	}

	results, err := s.PluginResultService.GetPluginResultsOrError(in.GetFileName(), userUUID)
	if err != nil {
		return &proto.GetPluginResultsResponse{}, err
	}

	protoList := make([]*proto.PluginResult, len(results))
	for i, res := range results {
		protoList[i] = ConvertPluginResultToProto(res)
	}

	return &proto.GetPluginResultsResponse{
		Results: protoList,
		Count:   uint32(len(protoList)),
	}, nil
}

func (s *PluginService) DeletePluginResult(ctx context.Context, in *proto.DeletePluginResultRequest) (*proto.DeletePluginResultResponse, error) {
	pluginResId := in.GetResultId()
	if len(pluginResId) == 0 {
		return &proto.DeletePluginResultResponse{}, errors.New("plugin result id is required")
	}
	err := s.PluginResultService.DeletePluginResult(pluginResId)
	if err != nil {
		return &proto.DeletePluginResultResponse{}, err
	}

	return &proto.DeletePluginResultResponse{}, nil
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
