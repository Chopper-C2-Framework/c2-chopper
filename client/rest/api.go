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
	app := fiber.New()

	app.Use(cors.New())

	app.Use(
		logger.New(),
	)

	// Uncomment when this is done
	// api := app.Group("/api")
	// routes.AuthRouter(api,AuthService)
	// routes.UserRouter(api, UserService)
	// routes.ManagementRouter(api, UserService)
	// routes.ReportRouter(api, UserService)
	// routes.AgentRouter(api,AgentService)

	err := app.Listen(fmt.Sprintf(":%d", frameworkConfig.ClientPort))

	if err != nil {
		return err
	}

	return nil
}
