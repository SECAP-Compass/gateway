package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"secap-gw/internal/server"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := server.New()
	server.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	server.RegisterFiberRoutes()
	server.RegisterBuildingRoutes()
	server.RegisterInventoryRoutes()

	p := os.Getenv("PORT")
	if p == "" {
		p = "5173"
	}

	port, _ := strconv.Atoi(p)
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
