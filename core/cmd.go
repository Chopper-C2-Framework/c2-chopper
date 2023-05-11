package core

import (
	"context"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"

	"github.com/urfave/cli/v2"
)

const (
	CLI_NAME = "c2-chopper"
)

func CreateApp(cmds ...[]*cli.Command) *cli.App {
	var commands []*cli.Command

	for _, moreCommands := range cmds {
		commands = append(commands, moreCommands...)
	}

	app := &cli.App{
		Name: CLI_NAME,
		Before: func(ctx *cli.Context) error {

			config := config.ParseConfigFromPath()
			ctx.Context = context.WithValue(ctx.Context, "config", config)

			dbConnection, _ := orm.CreateDB(config)
			ctx.Context = context.WithValue(ctx.Context, "dbConnection", dbConnection)
			return nil

		},
		Commands: commands,
		Authors: []*cli.Author{
			{
				Name:  "Yassine Belkhadem",
				Email: "yassine.belkhadem@insat.rnu.tn",
			},
			{
				Name:  "Mongi Saidane",
				Email: "mongi.saidane@insat.ucar.tn",
			},
		},
		Usage:                "A futuristic C2 framework where you can achieve all your dreams",
		Copyright:            "(c) 2023 Soter14",
		EnableBashCompletion: true,
	}

	return app
}
