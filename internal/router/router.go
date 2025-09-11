package router

import (
	//"tolelom_api/internal/handler"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// swagger 설정
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	app.Use(swagger.New(cfg))

	// API 라우트
	//api := app.Group("/api")
	//v1 := api.Group("/v1")

	//app.Get("/health", handler.HealthCheck)

}
