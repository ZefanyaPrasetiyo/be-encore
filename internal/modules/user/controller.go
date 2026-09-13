package user

import (
	"github.com/gofiber/fiber/v3"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetAll(ctx fiber.Ctx) error {
	users, err := c.service.GetAllUsers()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data pengguna",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
}