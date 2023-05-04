package config

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	generateConfigCommand := &cli.Command{
		Name:    "gen-config",
		Aliases: []string{"genc"},
		Usage:   "Generate default configuration for the framework.",

		Action: func(cCtx *cli.Context) error {
			err := GenerateConfigIfNotExists()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			return nil
		},
	}

	checkConfigFile := &cli.Command{
		Name:    "check-conf",
		Aliases: []string{"cconf"},
		Usage:   "Check if the configuration file is valid or not.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "default",
				Usage:   "Path to the configuration file",
			},
		},

		Action: func(cCtx *cli.Context) error {
			file := cCtx.String("file")
			if file == "default" {
				fmt.Fprintln(os.Stderr, "No configuration file was passed.")
			}

			config, err := ParseConfigFromFile(file)

			if err != nil {
				fmt.Fprintln(os.Stderr, "Error occured while parsing configuration file.", err)
				return nil
			}

			fmt.Println(config)

			return nil
		},
	}

	var commands []*cli.Command
	commands = append(commands, generateConfigCommand)
	commands = append(commands, checkConfigFile)

	return commands
}
