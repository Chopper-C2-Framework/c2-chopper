package main

import (
	"log"
	"os"

	"github.com/chopper-c2-framework/c2-chopper/core"
	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/core/plugins"

	server "github.com/chopper-c2-framework/c2-chopper/server"
)

func main() {

	configCommands := config.GetCommands()

	conf, err := config.GetConfig()
	if err != nil {
		log.Panicln("Meow")
	}

	plugins.CreatePluginManager(conf)
	serverCommands := server.GetCommands()

	framework := core.CreateApp(configCommands, serverCommands)
	// frameworkConfiguration := config.ParseConfigFromPath()

	err = framework.Run(os.Args)
	if err != nil {
		log.Panicln("Error occured while launching the framework", err)
	}
}
