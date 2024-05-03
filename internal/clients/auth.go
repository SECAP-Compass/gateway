package clients

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"

	"github.com/gofiber/fiber/v2"
)

type AuthClient interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Roles(ctx *fiber.Ctx) ([]string, error)
}

type authClient struct {
	baseUrl string
}

func NewAuthClient() AuthClient {
	baseUrl := os.Getenv("AUTH_BASEURL")
	if baseUrl == "" {
		panic("AUTH_BASEURL is not set")
	}

	return &authClient{
		baseUrl: baseUrl,
	}
}

func (c *authClient) Login(ctx *fiber.Ctx) error {

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

func (c *authClient) Register(ctx *fiber.Ctx) error {
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

func (c *authClient) Roles(ctx *fiber.Ctx) ([]string, error) {
	endpoint := fmt.Sprintf("%s/roles", c.baseUrl)
	a := fiber.Get(endpoint)

	a.Body(ctx.Body())
	a.Set("Content-Type", "application/json")

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return nil, ctx.Status(statusCode).JSON(
			fiber.Map{
				"errors": errs,
			},
		)
	}

	roles := make([]string, 0)
	if err := jsoniter.Unmarshal(body, roles); err != nil {
		return nil, ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Internal Server Error",
				"errors":  err,
			},
		)
	}

	return roles, nil
}
