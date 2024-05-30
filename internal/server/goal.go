package server

import "github.com/gofiber/fiber/v2"

func (s *FiberServer) RegisterGoalRoutes() {

	api := s.Group("/goals")
	api.Post("", s.AdminMiddleware, s.CreateGoalHandler)
}

func (s *FiberServer) CreateGoalHandler(c *fiber.Ctx) error {
	return s.inputClient.CreateGoal(c)
}
