package main

import (
	"log"
	"os"

	"github.com/chopper-c2-framework/c2-chopper/core"
	"github.com/chopper-c2-framework/c2-chopper/core/config"

	server "github.com/chopper-c2-framework/c2-chopper/server"
)

func setupCli() {

}
func x() error {
	return nil
}

func main() {

	configCommands := config.GetCommands()

	serverCommands := server.GetCommands()

	framework := core.CreateApp(configCommands, serverCommands)
	// frameworkConfiguration := config.ParseConfigFromPath()

	err := framework.Run(os.Args)
	if err != nil {
		log.Panicln("Error occured while launching the framework", err)
	}
}
