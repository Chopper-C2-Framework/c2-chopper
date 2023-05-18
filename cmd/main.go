package main

import (
	"log"
	"os"

	"github.com/chopper-c2-framework/c2-chopper/client"
	"github.com/chopper-c2-framework/c2-chopper/core"
	"github.com/chopper-c2-framework/c2-chopper/core/config"

	server "github.com/chopper-c2-framework/c2-chopper/server"
)

func main() {

	configCommands := config.GetCommands()

	serverCommands := server.GetCommands()
	clientCommands := client.GetCommands()

	framework := core.CreateApp(configCommands, serverCommands, clientCommands)

	err := framework.Run(os.Args)
	if err != nil {
		log.Panicln("Error occurred while launching the framework", err)
	}
}
