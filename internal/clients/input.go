package clients

import (
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
	//TODO implement me
	panic("implement me")
}

func (i *inputClient) MeasureBuilding(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
