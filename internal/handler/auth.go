package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func Authenticate(id, password string) error {
	const demoID = "admin"
	const demoPassword = "password"

	if (id == demoID) && (password == demoPassword) {
		return nil
	}
	return errors.New("invalid user")
}

func LoginHandler(c *fiber.Ctx) error {
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	if request.ID == "" || request.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "아이디와 비밀번호를 모두 입력해주세요.",
		})
	}

	if err := Authenticate(request.ID, request.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user",
		})
	}

	return c.JSON(LoginResponse{
		Message: "로그인 성공",
	})
}

func RegisterHandler(c *fiber.Ctx) error {
	var request RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	if request.ID == "" || request.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "아이디와 비밀번호를 모두 입력해주세요.",
		})
	}

	// TODO: 아이디 중복 검사, 비밀번호 해시 처리 등 실제 로직 추가 필요

	return c.JSON(fiber.Map{
		"message": "회원가입 성공",
	})
}
