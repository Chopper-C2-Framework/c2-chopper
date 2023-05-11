package routes

import (
	"github.com/chopper-c2-framework/c2-chopper/client/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service handlers.IAuthService) {
	app.Post("/login", handlers.Login(service))
	app.Post("/register", handlers.Register(service))
}
