package client

import (
	"github.com/chopper-c2-framework/c2-chopper/client/api"
	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {

	startServerCommand := &cli.Command{
		Name:    "client",
		Aliases: []string{"client"},
		Usage:   "Control the C2 client.",
		Action: func(cCtx *cli.Context) error {

			frameworkConfig := cCtx.Context.Value("config").(*config.Config)

			if frameworkConfig == nil {
				return nil
			}

			var clientManager api.ClientManager

			err := clientManager.LaunchClient(frameworkConfig)
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
