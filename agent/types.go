package main

import pb "github.com/chopper-c2-framework/c2-chopper/grpc/proto"

type Services struct {
	AgentService pb.AgentServiceClient
	TaskService  pb.TaskServiceClient
}

type AgentInfo struct {
	Uuid     string
	Username string
	UserId   string
	Hostname string
	Cwd      string
	HomeDir  string
}
