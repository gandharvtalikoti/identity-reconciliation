package routes

import (
	"github.com/gofiber/fiber/v2"
)

type IdentifyRequest struct {
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phoneNumber"`
}

func Identify(c *fiber.Ctx) error {
	var body IdentifyRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Later: Logic to handle identity
	return c.JSON(fiber.Map{"status": "ok", "data": body})
}
