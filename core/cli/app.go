package cli

import "github.com/urfave/cli/v2"

const (
	CLI_NAME = ""
)

func CreateApp(commands []*cli.Command) *cli.App {
	app := &cli.App{
		Name:     CLI_NAME,
		Commands: commands,
	}

	return app
}