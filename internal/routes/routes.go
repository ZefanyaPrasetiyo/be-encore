package routes

import (
	"encore-be/internal/modules/user"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	api := app.Group("api/v1")
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.

}