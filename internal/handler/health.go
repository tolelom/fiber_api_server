package handler

import "github.com/gofiber/fiber/v2"

type HealthResponse struct {
	Status string `json:"status" example:"ok"`
}

// HealthHandler godoc
//
//	@Summary		Health Check
//	@Description	Returns the health status of the server
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	HealthResponse
//	@Router			/health [get]
func HealthHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(HealthResponse{
		Status: "ok",
	})
}
