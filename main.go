package main

import (
	"Artify/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := api.SetUpRoute()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Artify API")
	})

	app.Listen(":8080")
}
