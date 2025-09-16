package model

type User struct {
	ID       string `gorm:"primaryKey;not null"`
	Password string `gorm:"not null"`
}
