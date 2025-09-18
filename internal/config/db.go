package config

import (
	"fmt"
	"log"
	"tolelom_api/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDataBase(cfg *Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("MySQL 연결에 실패했습니다: %v", err)
	}

	DB = db

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("DB instance에 연결 실패했습니다: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Printf("DB 핑 실패: %v", err)
	}

	log.Println("Database 연결 성공")

	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Printf("자동 마이그레이션 실패: %v", err)
		return err
	}

	log.Println("자동 마이그레이션 완료")

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
