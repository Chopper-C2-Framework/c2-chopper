package grpc

import (
	"context"
	"errors"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"github.com/google/uuid"

	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	services "github.com/chopper-c2-framework/c2-chopper/core/services"
)

type TaskService struct {
	proto.UnimplementedTaskServiceServer
	TaskService services.ITaskService
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

	// TODO: Add user id
	var task = entity.TaskModel{
		Name:    taskProto.GetName(),
		Args:    taskProto.GetArgs(),
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
