package handlers

import (
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"github.com/gofiber/fiber/v2"
)

type IAgentService proto.ListenerServiceClient

func GetAgents(agentService IAgentService) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Get Agents")

	}
}

func GetAgent(agentService IAgentService) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Get Agent")

	}
}

func SendCmd(agentService IAgentService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Send command")
	}
}

func GetResults(agentService IAgentService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Get results")
	}
}
