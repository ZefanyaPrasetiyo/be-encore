package user

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

func (h *Handler) GetAll(c fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *Handler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *Handler) Create(c fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.Create(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
