package clients

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type AuthClient struct {
	baseUrl string
}

func NewAuthClient() *AuthClient {
	baseUrl := os.Getenv("AUTH_BASEURL")
	if baseUrl == "" {
		panic("AUTH_BASEURL is not set")
	}

	return &AuthClient{
		baseUrl: baseUrl,
	}
}

func (c *AuthClient) Login(ctx *fiber.Ctx) error {

	endpoint := fmt.Sprintf("%s/login", c.baseUrl)

	a := fiber.Post(endpoint)

	a.Body(ctx.Body())
	a.Set("Content-Type", "application/json")
	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Internal Server Error",
				"errors":  errs,
			},
		)
	}

	return ctx.Status(statusCode).Send(body)
}

func (c *AuthClient) Register(ctx *fiber.Ctx) error {
	endpoint := fmt.Sprintf("%s/register", c.baseUrl)
	a := fiber.Post(endpoint)
	a.Debug()

	a.Body(ctx.Body())
	a.Set("Content-Type", "application/json")

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"errors": errs,
			},
		)
	}

	return ctx.Status(statusCode).Send(body)
}
