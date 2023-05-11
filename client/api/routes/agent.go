package routes

import (
	"github.com/chopper-c2-framework/c2-chopper/client/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func AgentRouter(app fiber.Router, service handlers.IAgentService) {
	app.Get("/agent", handlers.GetAgents(service))
	app.Get("/agent/:id", handlers.GetAgent(service))
	app.Post("/agent/:id", handlers.SendCmd(service))
	app.Get("/agent/:id/results", handlers.GetResults(service))
}
