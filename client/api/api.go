package api

import (
	"fmt"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func launchClient(frameworkConfig config.Config) {
	api := fiber.New()
	api.Use(cors.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api.Group("/api")

	api.Listen(fmt.Sprintf(":%d", frameworkConfig.ClientPort))
}
