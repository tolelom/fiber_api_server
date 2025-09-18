package handler

import (
	"errors"
	"strings"
	"time"
	"tolelom_api/internal/config"
	"tolelom_api/internal/model"
	"tolelom_api/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	validate *validator.Validate
	db       *gorm.DB
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		validate: validator.New(),
		db:       config.GetDB(),
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var request model.RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.validate.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	request.Name = strings.ToLower(strings.TrimSpace(request.Name))

	var existing model.User
	if err := h.db.Where("name = ?", request.Name).First(&existing).Error; err == nil {
		return fiber.NewError(fiber.StatusConflict, "User already exists")
	}

	hash, err := utils.HashPassword(request.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := model.User{
		Username: request.Name,
		Password: hash,
	}

	if err := h.db.Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create token")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   user.ToResponse(),
		"token":  token,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var request model.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.validate.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	request.Name = strings.ToLower(strings.TrimSpace(request.Name))

	var user model.User
	if err := h.db.Where("name = ?", request.Name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Database error")
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	user.LastLogin = time.Now()
	if err := h.db.Save(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update last login")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user":  user.ToResponse(),
			"token": token,
		},
	})
}
