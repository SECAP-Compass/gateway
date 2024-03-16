package server

import (
	"secap-gw/internal/clients"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App

	authClient *clients.AuthClient
}

func New() *FiberServer {

	authClient := clients.NewAuthClient()
	initOtlp()

	server := &FiberServer{
		App: fiber.New(),

		authClient: authClient,
	}

	return server
}
