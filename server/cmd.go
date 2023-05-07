package server

import (
	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/urfave/cli/v2"
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

			serverManager := gRPCServer{}
			err := serverManager.NewgRPCServer(frameworkConfig)
			if err != nil {
				return err
			}
			return nil
		},
	}

	var commands []*cli.Command

	commands = append(commands, startServerCommand)

	return commands
}
