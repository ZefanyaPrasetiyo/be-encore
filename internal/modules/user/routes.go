package user

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	repository := NewRepository(db)
	validator := NewValidator()
	service := NewService(repository, validator)
	handler := NewHandler(service)

	user := api.Group("/user")
	user.Get("/", handler.GetAll)
	user.Get("/:id", handler.GetByID)
	user.Post("/", handler.Create)

}