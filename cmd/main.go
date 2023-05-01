package main

import (
	"fmt"
	"log"

	"github.com/chopper-c2-framework/c2-chopper/core/plugins"
)

func main() {

	plugins, err := plugins.LoadPlugins()
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, plugin := range plugins {
		fmt.Println("[+]", plugin.Name)
	}

}
