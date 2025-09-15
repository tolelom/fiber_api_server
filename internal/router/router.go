package router

import (
	"tolelom_api/internal/handler"

	"github.com/gofiber/fiber/v2"

	_ "tolelom_api/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!\nI'm fiber Server!")
	})

	app.Get("/health", handler.HealthHandler)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
