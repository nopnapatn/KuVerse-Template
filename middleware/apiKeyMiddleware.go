package middleware

import (
	"kuverse/config"

	"github.com/gofiber/fiber/v2"
)

func ApiKeyMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("API-KEY")
	if apiKey != config.GetApiKey() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid API key"})
	}
	return c.Next()
}
