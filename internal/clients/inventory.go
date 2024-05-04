package clients

import "github.com/gofiber/fiber/v2"

type InventoryClient interface {
	GetCities(c *fiber.Ctx) error
	GetCityById(c *fiber.Ctx) error
}

type inventoryClient struct {
	baseUrl string
}

func NewInventoryClient() InventoryClient {
	return &inventoryClient{
		baseUrl: "http://localhost:8003",
	}
}

func (i *inventoryClient) GetCities(c *fiber.Ctx) error {
	const endpoint = "/cities"

	a := fiber.Get(i.baseUrl + endpoint)

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

func (i *inventoryClient) GetCityById(c *fiber.Ctx) error {
	const endpoint = "/cities/"

	a := fiber.Get(i.baseUrl + endpoint + c.Params("id"))
	a.Debug()

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
