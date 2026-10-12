package artist

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	repository := NewRepository(db)
	validator := NewValidator()
	service := NewService(repository, validator)
	handler := NewHandler(service)

	artists := api.Group("/artists")

	artists.Get("/", handler.GetAll)
	artists.Get("/:id", handler.GetByID)
	artists.Post("/", handler.Create)
	artists.Patch("/:id", handler.Update)
	artists.Delete("/:id", handler.Delete)
}
