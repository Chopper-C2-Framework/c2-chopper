package grpc

import (
	"errors"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
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
