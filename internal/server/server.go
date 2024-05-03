package server

import (
	"secap-gw/internal/clients"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App

	authClient  clients.AuthClient
	inputClient clients.InputClient
}

func New() *FiberServer {

	initOtlp()

	server := &FiberServer{
		App: fiber.New(),

		authClient:  clients.NewAuthClient(),
		inputClient: clients.NewInputClient(),
	}

	return server
}
