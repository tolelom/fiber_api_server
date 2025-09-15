package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
		// .env 파일 로드
		if err := godotenv.Load(); err != nil {
			log.Println(".env 파일을 찾을 수 없습니다.")
		}

		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		cfg = &Config{
			Port: port,
		}
	})

	return cfg
}
