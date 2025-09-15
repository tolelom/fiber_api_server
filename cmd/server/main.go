package main

import (
	"log"
	"tolelom_api/internal/config"
	"tolelom_api/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	app := fiber.New()

	router.Setup(app)

	addr := ":" + cfg.Port
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Fiber server failed to start: %v", err)
	}
}
