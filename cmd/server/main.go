package main

import (
	"log"
	"tolelom_api/internal/config"
	"tolelom_api/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("환경 변수 로딩 실패: %v", err)
	}

	config.InitDB()

	log.Println("MySQL DB 연결 성공")

	app := fiber.New()

	router.Setup(app)

	addr := ":" + cfg.Port
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Fiber server failed to start: %v", err)
	}
}
