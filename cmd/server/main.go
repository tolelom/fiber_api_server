package main

import (
	"log"
	"tolelom_api/internal/config"
	"tolelom_api/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	if err := config.InitDataBase(cfg); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	app := fiber.New()

	router.Setup(app)

	addr := ":" + cfg.Port
	log.Printf("Server listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Fiber server failed to start: %v", err)
	}
}
