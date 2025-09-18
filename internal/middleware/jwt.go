package middleware

import (
	"strings"
	"tolelom_api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		parts := strings.Fields(header)

		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or missing Authorization header")
		}

		claims, err := utils.ValidateJWT(parts[1])
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
		}

		c.Locals("username", claims.Username)
		return c.Next()
	}
}
