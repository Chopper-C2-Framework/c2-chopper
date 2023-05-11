package handlers

import (
	"github.com/chopper-c2-framework/c2-chopper/proto"
	"github.com/gofiber/fiber/v2"
)

type IAuthService proto.AuthServiceClient

func Login(authService IAuthService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.JSON("Login")

	}
}

func Register(authService IAuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Register")

	}
}
