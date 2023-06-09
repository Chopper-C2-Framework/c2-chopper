package server

import (
	"log"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/urfave/cli/v2"
)

// A channel to block the main thread. Well TODO: CHANGE IT IN A MORE PROPER WAY OF DOING THINGS
var (
	c = make(chan int)
)

func GetCommands() []*cli.Command {

	startServerCommand := &cli.Command{
		Name:    "server",
		Aliases: []string{"server"},
		Usage:   "Control the C2 server state and functionalities.",
		Action: func(cCtx *cli.Context) error {
			frameworkConfig := cCtx.Context.Value("config").(*config.Config)

			if frameworkConfig == nil {
				return nil
			}

			var serverManager IServerManager = &Manager{}

			go func() {
				err := serverManager.NewgRPCServer(frameworkConfig)
				if err != nil {
					log.Println("Error launching GRPC server")
					return
				}
			}()
			// if err != nil {
			// 	log.Panicln("Error launching server", err)
			// 	return err
			// }

			go func() {
				err := serverManager.NewgRPCServerHTTPGateway(frameworkConfig)
				if err != nil {
					log.Println("Error launching HTTP Gateway server")
					return
				}
			}()
			// if err != nil {
			// 	log.Panicln("Error while starting gRPC server HTTP gateway: ", err)
			// 	return nil
			// }
			c <- 1

			return nil
		},
	}

	var commands []*cli.Command

	commands = append(commands, startServerCommand)

	return commands
}
