package routes

import (
	"encore-be/internal/modules/user"
	"encore-be/internal/modules/venue"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	api := app.Group("api/v1")


	venue.RegisterRoutes(api, db)
	user.RegisterRoutes(api, db)
}