package server

import "github.com/gofiber/fiber/v2"

func (s *FiberServer) RegisterInventoryRoutes() {
	s.App.Get("/cities", s.GetCitiesHandler)
	s.App.Get("/cities/:id", s.GetCityByIdHandler)
}

func (s *FiberServer) GetCitiesHandler(c *fiber.Ctx) error {
	return s.inventoryClient.GetCities(c)
}

func (s *FiberServer) GetCityByIdHandler(c *fiber.Ctx) error {
	return s.inventoryClient.GetCityById(c)
}
