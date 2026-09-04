package network

import "github.com/gofiber/fiber/v3"

func NewApp() *fiber.App{
	app := fiber.New()

	api := app.Group("/api")

	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di API king jepan",

		})
	})
	return app
}