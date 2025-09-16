package router

import (
	"tolelom_api/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "tolelom_api/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World! I'm fiber Server!",
		})
	})

	app.Get("/health", handler.HealthHandler)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Post("/api/login", handler.LoginHandler)
}
