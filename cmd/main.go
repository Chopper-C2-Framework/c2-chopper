package main

import (
	"github.com/Chopper-C2-Framework/C2-Chopper/core/plugins"
	"fmt"
	"log"
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
