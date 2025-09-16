package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DSN  string
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

	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, errors.New("DB_DSN 환경변수 미설정")
	}

	return &Config{
		Port: port,
		DSN:  dsn,
	}, nil
}
