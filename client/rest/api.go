package rest

import (
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type IClientManager interface {
	launchClient(frameworkConfig config.Config) error
}

type ClientManager struct{}

func (c ClientManager) LaunchClient(frameworkConfig *config.Config) error {
	// Serve front from public file
	app := fiber.New()

	app.Use(cors.New())

	app.Use(
		logger.New(),
	)

	err := app.Listen(fmt.Sprintf(":%d", frameworkConfig.ClientPort))

	if err != nil {
		return err
	}

	return nil
}
