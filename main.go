package main

import (
	db "bitespeed-identity/database"
	"bitespeed-identity/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bitespeed Identity API is live 🚀")
	})

	// /identify route
	app.Post("/identify", routes.Identify)

	app.Listen(":3000")
}
