package main

import (
	"fmt"

	"github.com/c2-chopper/core/plugins"
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
