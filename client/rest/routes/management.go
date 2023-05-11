package routes

import (
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"github.com/gofiber/fiber/v2"
)

type ManagementService struct {
	TeamsService proto.TeamServiceClient
	// FindingService proto.FindingsServiceClient
	// ReportService proto.ReportServiceClient
}

func CreateTeam(managementService ManagementService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.JSON("CreateTeam")

	}
}

func DeleteTeam(managementService ManagementService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.JSON("DeleteTeam")

	}
}

func UpdateTeam(managementService ManagementService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.JSON("UpdateTeam")
	}
}
