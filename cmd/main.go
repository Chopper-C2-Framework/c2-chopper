package main

import (
	"fmt"

	"github.com/Chopper-C2-Framework/C2-Chopper/core/plugins"
)

func main() {

	plugins, err := plugins.LoadPlugins()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	for _, plugin := range plugins {
		fmt.Println("[+]", plugin.Name)

	}

}
