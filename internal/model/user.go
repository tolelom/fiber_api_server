package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"size:100;uniqueIndex;not null" validate:"required,min=2,max=20"`
	Password  string `gorm:"not null" validate:"required,min=8"`
	CreatedAt time.Time
	LastLogin time.Time
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponse struct {
	Status string   `json:"status"`
	Data   AuthData `json:"data"`
}

type AuthData struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		LastLogin: u.LastLogin,
	}
}
