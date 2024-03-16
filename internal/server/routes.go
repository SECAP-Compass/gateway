package server

import (
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/propagation"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.Use(otelfiber.Middleware(otelfiber.WithPropagators(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))))
	s.Use(correlationIdMiddleware)
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Post("/login", s.LoginHandler)
	s.App.Post("/register", s.RegisterHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) LoginHandler(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	return s.authClient.Login(c)
}

func (s *FiberServer) RegisterHandler(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	return s.authClient.Register(c)
}
