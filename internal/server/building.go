package server

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const buildingAdminRole = "buildingAdmin"

func (s *FiberServer) RegisterBuildingRoutes() {
	s.App.Post("/buildings", s.AdminMiddleware, s.CreateBuildingHandler)

	// middleware?
	s.App.Get("/buildings", s.GetBuildingsHandler)
	s.App.Get("/buildings/filter", s.GetBuildingFilter)
	s.App.Get("/buildings/measurement-types", s.GetMeasurementTypeHeaders)
	s.App.Get("/buildings/measurement-types/:header", s.GetMeasurementTypes)
	s.App.Get("/buildings/measurement-units", s.GetMeasurementUnits)
	s.App.Get("/buildings/:id", s.GetBuildingByIdHandler)
	s.App.Post("/buildings/:id/measurements", s.AdminMiddleware, s.MeasureBuildingHandler)
	s.App.Get("/buildings/:id/measurements", s.AdminMiddleware, s.GetBuildingMeasurementsById)

}

var measurementUnits = []string{
	"GJ", "MJ", "MWh", "KWh", "therm", "mmBtu", "Btu",
}

// There should expire control
func (s *FiberServer) AdminMiddleware(c *fiber.Ctx) error {
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
	c.Locals("X-Agent", claims.StandardClaims.Subject)
	return c.Next()
}

func (s *FiberServer) CreateBuildingHandler(c *fiber.Ctx) error {
	return s.inputClient.CreateBuilding(c)
}

func (s *FiberServer) MeasureBuildingHandler(c *fiber.Ctx) error {
	return s.inputClient.MeasureBuilding(c, c.Params("id"))
}

func (s *FiberServer) GetBuildingFilter(c *fiber.Ctx) error {
	return s.inventoryClient.GetBuildingFilter(c)
}

func (s *FiberServer) GetBuildingMeasurementsById(c *fiber.Ctx) error {
	return s.inventoryClient.GetBuildingMeasurementsById(c)
}

func (s *FiberServer) GetMeasurementTypeHeaders(c *fiber.Ctx) error {
	return s.inputClient.GetMeasurementTypeHeaders(c)
}

func (s *FiberServer) GetMeasurementTypes(c *fiber.Ctx) error {
	return s.inputClient.GetMeasurementTypes(c)
}

func (s *FiberServer) GetBuildingsHandler(c *fiber.Ctx) error {
	return s.inventoryClient.GetBuildings(c)
}

func (s *FiberServer) GetBuildingByIdHandler(c *fiber.Ctx) error {
	return s.inventoryClient.GetBuildingById(c)

}

func (s *FiberServer) GetMeasurementUnits(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"measurementUnits": measurementUnits,
	})
}
