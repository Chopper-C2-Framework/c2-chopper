package grpc

import (
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"

	"github.com/mattn/go-shellwords"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func ConvertTeamToProto(team *entity.TeamModel) *proto.Team {
	var users []*proto.User
	for _, m := range team.Members {
		users = append(users, ConvertUserToProto(&m))
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

func ConvertProtoToUser(user *proto.User) (*entity.UserModel, error) {
	parsedUuid, err := uuid.Parse(user.Id)
	if err != nil {
		log.Debugf("ConvertProtoToUser: error parsing uuid %v\n", err)
		return nil, err
	}
	return &entity.UserModel{
		UUIDModel: entity.UUIDModel{ID: parsedUuid},
		Username:  user.Username,
	}, nil
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

func ConvertPluginResultToProto(res *entity.PluginResultModel) *proto.PluginResult {
	return &proto.PluginResult{
		Id:         res.ID.String(),
		Path:       res.Path,
		Output:     res.Output,
		OutputType: res.OutputType,
		CreatedAt:  res.CreatedAt.String(),
	}
}

func ConvertTaskTypeToProto(task *entity.TaskModel) proto.TaskType {
	if task.Type == entity.TASK_TYPE_PING {
		return proto.TaskType_PING
	}
	if task.Type == entity.TASK_TYPE_SHELL {
		return proto.TaskType_SHELL
	}
	return proto.TaskType_UNKNOWN
}

func ConvertTaskToProto(task *entity.TaskModel) *proto.Task {
	var args []string
	if task.Type == entity.TASK_TYPE_SHELL {
		args, _ = shellwords.Parse(task.Args)
	} else {
		args = make([]string, 1)
		args[0] = task.Args
	}
	// TODO: Add user id
	return &proto.Task{
		TaskId:  task.ID.String(),
		Name:    task.Name,
		Args:    args,
		Type:    ConvertTaskTypeToProto(task),
		AgentId: task.AgentId.String(),
	}
}

func ConvertAgentToProto(agent *entity.AgentModel) *proto.Agent {
	return &proto.Agent{
		Id:        agent.ID.String(),
		Nickname:  agent.Nickname,
		Hostname:  agent.Hostname,
		Username:  agent.Username,
		UserId:    agent.Uid,
		SleepTime: agent.SleepTime,
	}
}

func ConvertTaskResultToProto(taskResult *entity.TaskResultModel) *proto.TaskResult {
	return &proto.TaskResult{
		Id:     taskResult.ID.String(),
		Status: taskResult.Status,
		TaskId: taskResult.TaskID.String(),
		Output: taskResult.Output,
		Seen:   taskResult.Seen,
	}
}
