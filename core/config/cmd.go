package config

import "github.com/urfave/cli/v2"

func GetCommands() []*cli.Command {
	generateConfigCommand := &cli.Command{
		Name:    "gen-config",
		Aliases: []string{"genc"},
		Usage:   "Generate default configuration for the framework.",
		Action: func(cCtx *cli.Context) error {

			return nil
		},
	}

	checkConfigFile := &cli.Command{
		Name:    "check-conf",
		Aliases: []string{"cconf"},
		Usage:   "Check if the configuration file is valid or not.",
		Action: func(cCtx *cli.Context) error {

			return nil
		},
	}

	var commands []*cli.Command
	commands = append(commands, generateConfigCommand)
	commands = append(commands, checkConfigFile)

	return commands
}
