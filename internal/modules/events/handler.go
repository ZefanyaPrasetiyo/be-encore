package events

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func handleError(c fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "event tidak ditemukan",
		})
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validasi request gagal",
			"errors":  validationErrors.Error(),
		})
	}

	message := err.Error()

	if message == "artist tidak ditemukan" ||
		message == "venue tidak ditemukan" ||
		message == "minimal satu field harus diisi" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": message,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "terjadi kesalahan pada server",
	})
}

func (h *Handler) GetAll(c fiber.Ctx) error {
	events, err := h.service.GetAllEvents()
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengambil data event",
		"data":    ToResponses(events),
	})
}

func (h *Handler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "format UUID tidak valid",
		})
	}

	event, err := h.service.GetEventById(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengambil detail event",
		"data":    ToResponse(*event),
	})
}

func (h *Handler) Create(c fiber.Ctx) error {
	var req CreateEventRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "request body tidak valid",
		})
	}

	event, err := h.service.CreateEvent(req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "event berhasil dibuat",
		"data":    ToResponse(*event),
	})
}

func (h *Handler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "format UUID tidak valid",
		})
	}

	var req UpdateEventRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "request body tidak valid",
		})
	}

	event, err := h.service.UpdateEvent(id, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "event berhasil diperbarui",
		"data":    ToResponse(*event),
	})
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "format UUID tidak valid",
		})
	}

	if err := h.service.DeleteEvent(id); err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "event berhasil dihapus",
	})
}