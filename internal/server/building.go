package server

import (
	"github.com/gofiber/fiber/v2"
	"slices"
	"strings"
)

const buildingAdminRole = "buildingAdmin"

func (s *FiberServer) RegisterBuildingRoutes() {
	s.App.Post("/building", s.buildingAdminMiddleware, s.CreateBuildingHandler)
	s.App.Get("/building/measure", s.buildingAdminMiddleware, s.MeasureBuildingHandler)
}

func (s *FiberServer) buildingAdminMiddleware(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")
	if bearerToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	token := strings.Split(bearerToken, "Bearer ")[1]
	claims, err := s.jwtHandler.FetchClaims(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !slices.Contains(claims.Roles, buildingAdminRole) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	c.Locals("X-Authority", claims.Authority)
	return c.Next()
}

func (s *FiberServer) CreateBuildingHandler(c *fiber.Ctx) error {
	return s.inputClient.CreateBuilding(c)
}

func (s *FiberServer) MeasureBuildingHandler(c *fiber.Ctx) error {
	return s.inputClient.MeasureBuilding(c)
}
