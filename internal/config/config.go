package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		log.Println(".env 파일을 찾을 수 없습니다.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port: port,
	}, nil
}
