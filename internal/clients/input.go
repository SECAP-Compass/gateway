package clients

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

type InputClient interface {
	CreateBuilding(c *fiber.Ctx) error
	MeasureBuilding(c *fiber.Ctx) error
	GetMeasurementTypeHeaders(c *fiber.Ctx) error
	GetMeasurementTypes(c *fiber.Ctx) error
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
	a.Set("X-Authority", c.Locals("X-Authority").(string))
	a.Set("X-Agent", c.Locals("X-Agent").(string))

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

	a.Body(c.Body())
	a.Set("Content-Type", "application/json")
	a.Set("X-Authority", c.Get("X-Authority"))

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

func (i *inputClient) GetMeasurementTypeHeaders(c *fiber.Ctx) error {
	endpoint := fmt.Sprintf("%s/building/measurement-types", i.baseUrl)

	a := fiber.Get(endpoint)

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return c.Status(statusCode).JSON(
			fiber.Map{
				"errors": errs,
			},
		)
	}

	c.Set("Content-Type", "application/json")
	return c.Status(statusCode).Send(body)
}

func (i *inputClient) GetMeasurementTypes(c *fiber.Ctx) error {
	endpoint := fmt.Sprintf("%s/building/measurement-types/%s", i.baseUrl, c.Params("header"))

	a := fiber.Get(endpoint)

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return c.Status(statusCode).JSON(
			fiber.Map{
				"errors": errs,
			},
		)
	}

	c.Set("Content-Type", "application/json")
	return c.Status(statusCode).Send(body)
}
