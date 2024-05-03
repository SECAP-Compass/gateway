package server

import (
	"github.com/gofiber/fiber/v2"
	"slices"
)

func RegisterBuildingRoutes(s *FiberServer) {
	s.App.Post("/building", s.buildingAdminMiddleware, s.CreateBuildingHandler)
	s.App.Get("/building/measure", s.buildingAdminMiddleware, s.MeasureBuildingHandler)
}

func (s *FiberServer) buildingAdminMiddleware(c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	roles, err := s.authClient.Roles(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	if !slices.Contains(roles, "buildingAdmin") {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	return c.Next()
}

func (s *FiberServer) CreateBuildingHandler(c *fiber.Ctx) error {

}

func (s *FiberServer) MeasureBuildingHandler(c *fiber.Ctx) error {

}
