package clients

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

type InputClient interface {
	CreateBuilding(c *fiber.Ctx) error
	MeasureBuilding(c *fiber.Ctx) error
}

type inputClient struct {
	baseUrl string
}

func NewInputClient() InputClient {
	baseUrl := os.Getenv("INPUT_BASEURL")
	if baseUrl == "" {
		panic("INPUT_BASEURL is not set")
	}

	return &inputClient{
		baseUrl: baseUrl,
	}
}

func (i *inputClient) CreateBuilding(c *fiber.Ctx) error {
	endpoint := fmt.Sprintf("%s/building", i.baseUrl)

	a := fiber.Post(endpoint)
	a.Debug()

	a.Body(c.Body())
	a.Set("Content-Type", "application/json")

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return c.Status(statusCode).JSON(
			fiber.Map{
				"errors": errs,
			},
		)
	}

	return c.Status(statusCode).Send(body)
}

func (i *inputClient) MeasureBuilding(c *fiber.Ctx) error {

	endpoint := fmt.Sprintf("%s/building/measure", i.baseUrl)

	a := fiber.Post(endpoint)
	a.Debug()

	a.Body(c.Body())
	a.Set("Content-Type", "application/json")

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return c.Status(statusCode).JSON(
			fiber.Map{
				"errors": errs,
			},
		)
	}

	return c.Status(statusCode).Send(body)
}
