package main

import (
	"fmt"
	"log"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
	"github.com/chopper-c2-framework/c2-chopper/server/grpc"
)

func main() {

	config, err := config.GetConfig()

	if err != nil {
		log.Fatalln("Error getting configuration file, if it's the first time launching the framework add --init")
	}

	plugins, err := plugins.LoadPlugins()

	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, plugin := range plugins {
		fmt.Println("[+]", plugin.Name)
	}

	grpc.NewgRPCServer(*config)
}
