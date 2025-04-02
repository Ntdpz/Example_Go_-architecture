package handlers

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/internal/services"

)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUserHandler(c *fiber.Ctx) error {
	users, err := h.Service.GetUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}
	return c.JSON(fiber.Map{
		"count": len(users),
		"data":  users,
	})

}
