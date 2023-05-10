package api

import (
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type IClientManager interface {
	launchClient(frameworkConfig config.Config) error
}

type ClientManager struct{}

func (c ClientManager) LaunchClient(frameworkConfig *config.Config) error {
	api := fiber.New()
	api.Use(cors.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api.Group("/api")

	err := api.Listen(fmt.Sprintf(":%d", frameworkConfig.ClientPort))

	if err != nil {
		return err
	}

	return nil
}
