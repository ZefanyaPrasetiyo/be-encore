package events

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c fiber.Ctx) error {
	events, err := h.service.GetAllEvents()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengambil data event",
		"data":    ToResponses(events),
	})
}

func (h *Handler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "format UUID tidak valid")
	}

	event, err := h.service.GetEventById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengambil detail event",
		"data":    ToResponse(*event),
	})
}

func (h *Handler) Create(c fiber.Ctx) error {
	var req CreateEventRequest

	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "request body tidak valid")
	}

	event, err := h.service.CreateEvent(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "event berhasil dibuat",
		"data":    ToResponse(*event),
	})
}

func (h *Handler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "format UUID tidak valid")
	}

	var req UpdateEventRequest

	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "request body tidak valid")
	}

	event, err := h.service.UpdateEvent(id, req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "event berhasil diperbarui",
		"data":    ToResponse(*event),
	})
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "format UUID tidak valid")
	}

	if err := h.service.DeleteEvent(id); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "event berhasil dihapus",
	})
}