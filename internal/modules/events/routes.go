package events

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	repository := NewRepository(db)
	validator := NewValidator()
	service := NewService(repository, validator)
	handler := NewHandler(service)

	events := api.Group("/events")

	events.Get("/", handler.GetAll)
	events.Get("/:id", handler.GetByID)
	events.Post("/", handler.Create)
	events.Patch("/:id", handler.Update)
	events.Delete("/:id", handler.Delete)
}