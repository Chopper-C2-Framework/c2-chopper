package grpc

import (
	"errors"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"github.com/google/uuid"
)

func ValidateTaskProto(task *proto.Task) error {
	if task == nil {
		return errors.New("Task info are required")
	}
	if len(task.GetName()) == 0 {
		return errors.New("Task name is required")
	}
	taskType := task.GetType()
	if taskType == proto.TaskType_UNKNOWN {
		return errors.New("Task type is required")
	}
	return nil
}

func ValidateTaskResultProto(taskResult *proto.TaskResult) error {
	if taskResult == nil {
		return errors.New("task result is required.")
	}
	if len(taskResult.GetTaskId()) == 0 {
		return errors.New("Task id is required.")
	}
	_, err := uuid.Parse(taskResult.GetTaskId())
	if err != nil {
		return errors.New("Invalid task id")
	}
	return nil
}
