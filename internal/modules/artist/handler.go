package artist

import (
	"errors"

	"github.com/gofiber/fiber/v3"
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
			"message": "data tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func (h *Handler) GetAll(c fiber.Ctx) error {
	artists, err := h.service.GetAllArtist()
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengambil data artist",
		"data":    ToResponses(artists),
	})
}

func (h *Handler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "format UUID tidak valid",
		})
	}

	artist, err := h.service.GetArtistById(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengambil data artist",
		"data":    ToResponse(*artist),
	})
}

func (h *Handler) Create(c fiber.Ctx) error {
	var req CreateArtistRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "request body tidak valid",
		})
	}

	artist, err := h.service.CreateArtist(req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "artist berhasil dibuat",
		"data":    ToResponse(*artist),
	})
}

func (h *Handler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "format UUID tidak valid",
		})
	}

	var req UpdateArtistRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "request body tidak valid",
		})
	}

	artist, err := h.service.UpdateArtist(id, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "artist berhasil diperbarui",
		"data":    ToResponse(*artist),
	})
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "format UUID tidak valid",
		})
	}

	if err := h.service.DeleteArtist(id); err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "artist berhasil dihapus",
	})
}
