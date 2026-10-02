package venue

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	repository := NewRepository(db)
	validator := NewValidator()
	service := NewService(repository, validator)
	handler := NewHandler(service)

	category := api.Group("/venue-category")
	category.Get("/", handler.GetVenueCategory)
	category.Get("/:id", handler.GetVenueCategoryById)
	category.Post("/", handler.CreateVenueCategory)
	category.Put("/:id", handler.UpdateVenueCategory)

	venue := api.Group("/venue")
	venue.Get("/", handler.GetVenue)
	venue.Get("/:id", handler.GetVenueById)
	venue.Post("/", handler.CreateVenue)
}