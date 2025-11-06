package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unomcp/JueJin-MCP/middleware"
)

func start(port string) {
	app := newApp(fiber.New(fiber.Config{
		DisableStartupMessage: false,
	}))

	app.FiberApp.Use(middleware.Cros())

	setupRoutes(app)

	log.Fatal(app.FiberApp.Listen(port))
}

func main() {
	start(":10086")
}
