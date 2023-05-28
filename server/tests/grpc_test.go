package server_test

import (
	"fmt"
	"testing"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCServices(t *testing.T) {
	defaultConfig := config.CreateDefaultConfig()

	var serverManager server.IServerManager = &server.Manager{}

	go func() {
		err := serverManager.NewgRPCServer(defaultConfig)
		if err != nil {
			t.Errorf("error starting gRPC Server: %v", err)
			return
		}
	}()

	host := fmt.Sprintf("%s:%d", defaultConfig.Host, defaultConfig.ServergRPCPort)
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("error connecting to gRPC Server: %v", err)
		return
	}

	t.Run("PluginService", func(subTest *testing.T) {
		gRPCPluginService_Test(subTest, conn)
		fmt.Println("PluginService is done")
	})

	t.Cleanup(func() {
		fmt.Println("CLOSING SERVER")
		serverManager.CloseGRPCServer()
	})
}
