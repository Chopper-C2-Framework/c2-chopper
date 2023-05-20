package grpc

import (
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

func ConvertTeamToProto(team *entity.TeamModel) *proto.Team {
	var users []*proto.User
	for _, m := range team.Members {
		users = append(users, ConvertUserToProto(m))
	}
	protoTeam := &proto.Team{
		Id:      team.ID.String(),
		Name:    team.Name,
		Members: users,
	}

	return protoTeam
}

func ConvertUserToProto(user *entity.UserModel) *proto.User {

	return &proto.User{
		Username: user.Username,
		Id:       user.ID.String(),
	}
}

func ConvertPluginToProto(plugin plugins.IPlugin) *proto.Plugin {
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
