package server

import "github.com/urfave/cli/v2"

func GetCommands(startServer func() error) []*cli.Command {
	startServerCommand := &cli.Command{
		Name:    "server",
		Aliases: []string{"server"},
		Usage:   "Control the C2 server state and functionalities.",
		Action: func(cCtx *cli.Context) error {
			err := startServer()
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
