package venue

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetVenueCategory(c fiber.Ctx) error{
	category, err := h.service.GetAllVenueCategory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve venue categories",
		})
	}

	return c.JSON(category)
}

func (h *Handler) GetVenueCategoryById(c fiber.Ctx) error{
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid venue category ID format",
		})
	}

	category, err := h.service.GetVenueCategoryById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Venue category not found",
		})
	}

	return c.JSON(category)
}

func (h *Handler) CreateVenueCategory(c fiber.Ctx) error{
	var req CreateVenueCategoryRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	category, err := h.service.CreateVenueCategory(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create venue category",
		})
	}

	return c.JSON(category)
}

func (h *Handler) UpdateVenueCategory(c fiber.Ctx) error{

	var req UpdateVenueCategoryRequest

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid venue category ID format",
		})
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	_, err = h.service.GetVenueCategoryById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Venue category not found",
		})
	}

	category, err := h.service.UpdateVenueCategory(id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update venue category",
		})
	}

	return c.JSON(category)
}


func (h *Handler) DeleteVenueCategory(c fiber.Ctx) error{
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid venue category ID format",
		})
	}

	_, err = h.service.GetVenueCategoryById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Venue category not found",
		})
	}

	err = h.service.DeleteVenueCategory(id)
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"error": "Failed to delete venue category",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Venue category deleted successfully",
	})
}

func (h *Handler) GetVenue(c fiber.Ctx) error{
	venue, err := h.service.GetAllVenue()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve venues",
		})
	}

	return c.JSON(venue)
}

func (h *Handler) GetVenueById(c fiber.Ctx) error{
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid venue ID format",
		})	
	}

	category, err := h.service.GetVenueById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Venue not found",
		})
	}
	
	return c.JSON(category)
}

func (h *Handler) CreateVenue(c fiber.Ctx)error{
	var req CreateVenueRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	venue, err :=  h.service.CreateVenue(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"Failed to create venue",
		})
	}

	return c.JSON(venue)
}

func (h *Handler) UpdateVenue(c fiber.Ctx) error{
	var req UpdateVenueRequest

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid venue ID format",
		})

	}
	if err :=  c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	_, err = h.service.GetVenueById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Venue not found",
		})
	}

	venue, err := h.service.updateVenue(id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update venue",
		})
	}

	return c.JSON(venue)
}

func (h *Handler) DeleteVenue(c fiber.Ctx) error{
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid venue ID format",
		})
	}

	_, err = h.service.GetVenueById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Venue not found",
		})
	}

	err = h.service.DeleteVenue(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"failed to delete venue",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Venue deleted successfully",
	})
}


