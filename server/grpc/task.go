package grpc

import (
	"context"
	"errors"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
	services "github.com/chopper-c2-framework/c2-chopper/core/services"
)

type TaskService struct {
	proto.UnimplementedTaskServiceServer
	TaskService   services.ITaskService
	PluginManager plugins.IPluginManager
	AgentService  services.IAgentService
}

func (s *TaskService) GetTask(ctx context.Context, in *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
	if len(in.GetTaskId()) == 0 {
		return &proto.GetTaskResponse{}, errors.New("Task id required")
	}

	task, err := s.TaskService.FindTaskOrError(in.GetTaskId())
	if err != nil {
		return &proto.GetTaskResponse{}, errors.New("Task not found")
	}
	return &proto.GetTaskResponse{
		Task: ConvertTaskToProto(task),
	}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, in *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	if len(in.GetTaskId()) == 0 {
		return &proto.DeleteTaskResponse{}, errors.New("Task id required")
	}

	task, err := s.TaskService.FindTaskOrError(in.GetTaskId())
	if err != nil {
		return &proto.DeleteTaskResponse{}, errors.New("Task not found")
	}

	s.TaskService.DeleteTask(task)
	return &proto.DeleteTaskResponse{}, nil
}

func (s *TaskService) CreateTask(ctx context.Context, in *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	if len(in.GetAgentId()) == 0 {
		return &proto.CreateTaskResponse{}, errors.New("Agent id required")
	}
	agentId, err := uuid.Parse(in.GetAgentId())
	if err != nil {
		return &proto.CreateTaskResponse{}, errors.New("Invalid agent id")
	}

	taskProto := in.GetTask()
	err = ValidateTaskProto(taskProto)
	if err != nil {
		return &proto.CreateTaskResponse{}, err
	}

	args := ""
	if len(taskProto.GetArgs()) != 0 {
		args = taskProto.GetArgs()[0]
	}
	// TODO: Add user id
	var task = entity.TaskModel{
		Name:    taskProto.GetName(),
		Args:    args,
		Type:    entity.TaskType(taskProto.GetType().String()),
		AgentId: agentId,
		// CreatorId: ,
	}

	err = s.TaskService.CreateTask(&task)
	if err != nil {
		return &proto.CreateTaskResponse{}, err
	}

	return &proto.CreateTaskResponse{}, nil
}

func (s *TaskService) GetAllTasks(ctx context.Context, in *emptypb.Empty) (*proto.GetAllTasksResponse, error) {
	tasks, err := s.TaskService.FindAllTasks()
	if err != nil {
		return &proto.GetAllTasksResponse{}, err
	}

	protoList := make([]*proto.Task, len(tasks))
	for i, task := range tasks {
		protoList[i] = ConvertTaskToProto(task)
	}

	return &proto.GetAllTasksResponse{
		Tasks: protoList,
		Count: uint32(len(protoList)),
	}, nil
}

func (s *TaskService) GetAgentTasks(ctx context.Context, in *proto.GetAgentTasksRequest) (*proto.GetAgentTasksResponse, error) {
	agentId := in.GetAgentId()
	if len(agentId) == 0 {
		return &proto.GetAgentTasksResponse{}, errors.New("Agent id required")
	}

	agent, err := s.AgentService.FindAgentOrError(agentId)
	if err != nil {
		return &proto.GetAgentTasksResponse{}, errors.New("Agent not found")
	}

	tasks, err := s.TaskService.FindTasksForAgent(agentId)
	if err != nil {
		return &proto.GetAgentTasksResponse{}, err
	}

	protoList := make([]*proto.Task, len(tasks))
	for i, task := range tasks {
		protoList[i] = ConvertTaskToProto(task)
	}

	return &proto.GetAgentTasksResponse{
		Tasks:     protoList,
		SleepTime: agent.SleepTime,
	}, nil
}

func (s *TaskService) GetAgentUnexecutedTasks(ctx context.Context, in *proto.GetAgentUnexecutedTasksRequest) (*proto.GetAgentUnexecutedTasksResponse, error) {
	agentId := in.GetAgentId()
	if len(agentId) == 0 {
		return &proto.GetAgentUnexecutedTasksResponse{}, errors.New("Agent id required")
	}

	agent, err := s.AgentService.FindAgentOrError(agentId)
	if err != nil {
		return &proto.GetAgentUnexecutedTasksResponse{}, errors.New("Agent not found")
	}

	tasks, err := s.TaskService.FindUnexecutedTasksForAgent(agentId)
	if err != nil {
		return &proto.GetAgentUnexecutedTasksResponse{}, err
	}

	protoList := make([]*proto.Task, len(tasks))
	for i, task := range tasks {
		protoList[i] = ConvertTaskToProto(task)
	}

	return &proto.GetAgentUnexecutedTasksResponse{
		Tasks:     protoList,
		SleepTime: agent.SleepTime,
	}, nil
}

func (s *TaskService) GetActiveTasks(ctx context.Context, in *emptypb.Empty) (*proto.GetActiveTasksResponse, error) {
	tasks, err := s.TaskService.FindUnexecutedTasks()
	if err != nil {
		return &proto.GetActiveTasksResponse{}, err
	}

	protoList := make([]*proto.Task, len(tasks))
	for i, task := range tasks {
		protoList[i] = ConvertTaskToProto(task)
	}

	return &proto.GetActiveTasksResponse{
		Tasks: protoList,
		Count: uint32(len(protoList)),
	}, nil
}

func (s *TaskService) CreateTaskResult(ctx context.Context, in *proto.CreateTaskResultRequest) (*proto.CreateTaskResultResponse, error) {
	agentInfo := in.GetInfo()
	if agentInfo != nil {
		s.AgentService.ConnectAgent(
			agentInfo.Id,
			&entity.AgentModel{
				Username: agentInfo.GetUsername(),
				Uid:      agentInfo.GetUserId(),
				Hostname: agentInfo.GetHostname(),
				Cwd:      agentInfo.GetCwd(),
			},
		)
	}

	taskResProto := in.GetTaskResult()
	err := ValidateTaskResultProto(taskResProto)
	if err != nil {
		return &proto.CreateTaskResultResponse{}, err
	}

	taskUUID := uuid.MustParse(taskResProto.GetTaskId())
	taskResult := &entity.TaskResultModel{
		Status: taskResProto.GetStatus(),
		Output: taskResProto.GetOutput(),
		TaskID: taskUUID,
	}

	err = s.TaskService.CreateTaskResult(taskResult)
	if err != nil {
		return &proto.CreateTaskResultResponse{}, err
	}

	plugins := s.PluginManager.ListLoadedPlugins()
	for _, plugin := range plugins {
		loadedPlugin, err := s.PluginManager.GetPlugin(plugin)
		if err != nil {
			continue
		}
		if loadedPlugin.Channel == nil {
			continue
		}
		waiting, taskId := loadedPlugin.Plugin.IsWaitingForTaskResult()
		if waiting == false {
			continue
		}
		if taskId != taskUUID.String() {
			continue
		}
		loadedPlugin.Channel <- taskResult
	}

	return &proto.CreateTaskResultResponse{}, nil
}

func (s *TaskService) GetTaskResults(ctx context.Context, in *proto.GetTaskResultsRequest) (*proto.GetTaskResultsResponse, error) {
	if len(in.GetTaskId()) == 0 {
		return &proto.GetTaskResultsResponse{}, errors.New("Task id required")
	}

	taskResults, err := s.TaskService.FindTaskResults(in.GetTaskId())
	if err != nil {
		return &proto.GetTaskResultsResponse{}, err
	}

	protoList := make([]*proto.TaskResult, len(taskResults))
	for i, taskRes := range taskResults {
		protoList[i] = ConvertTaskResultToProto(taskRes)
	}

	return &proto.GetTaskResultsResponse{
		Results: protoList,
	}, nil
}

func (s *TaskService) GetLatestTaskResults(ctx context.Context, in *proto.GetLatestTaskResultsRequest) (*proto.GetLatestTaskResultsResponse, error) {
	limit := in.GetLimit()
	if limit == 0 {
		limit = 10
	}
	page := in.GetPage()
	if page == 0 {
		page = 1
	}

	var (
		taskResults []*entity.TaskResultModel
		err         error
	)
	if in.GetUnseen() {
		taskResults, err = s.TaskService.FindLatestUnseenResults(limit, page-1)
	} else {
		taskResults, err = s.TaskService.FindLatestResults(limit, page-1)
	}

	if err != nil {
		return &proto.GetLatestTaskResultsResponse{}, err
	}

	protoList := make([]*proto.TaskResult, len(taskResults))
	for i, taskRes := range taskResults {
		protoList[i] = ConvertTaskResultToProto(taskRes)
	}

	return &proto.GetLatestTaskResultsResponse{
		Results: protoList,
		Count:   uint32(len(protoList)),
	}, nil
}

func (s *TaskService) SetTaskResultsSeen(ctx context.Context, in *proto.SetTaskResultsSeenRequest) (*proto.SetTaskResultsSeenResponse, error) {
	resultIds := in.GetResultIds()
	if resultIds == nil {
		return &proto.SetTaskResultsSeenResponse{}, errors.New("At least 1 id is required in result_ids")
	}

	for _, id := range resultIds {
		err := s.TaskService.MarkTaskResultSeen(id)
		if err != nil {
			return &proto.SetTaskResultsSeenResponse{}, err
		}
	}
	return &proto.SetTaskResultsSeenResponse{}, nil
}
