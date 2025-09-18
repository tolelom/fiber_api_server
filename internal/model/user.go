package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null" validate:"required, min=2, max=20"`
	Password  string    `json:"-" gorm:"not null" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}
type LoginRequest struct {
	Name     string `json:"name" validate:"required,min=2, max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2, max=20"`
	Password string `json:"password" validate:"required,min=8"`
}
