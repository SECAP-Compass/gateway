package server

import (
	"secap-gw/internal/clients"
	"secap-gw/internal/jwt"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App

	authClient  clients.AuthClient
	inputClient clients.InputClient

	jwtHandler jwt.Handler
}

func New() *FiberServer {

	initOtlp()

	server := &FiberServer{
		App: fiber.New(),

		authClient:  clients.NewAuthClient(),
		inputClient: clients.NewInputClient(),
		jwtHandler:  jwt.NewJwtHandler(),
	}

	return server
}
