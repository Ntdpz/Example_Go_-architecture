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

	username := c.Query("Username")
	token := c.Query("Token")

	if username != "" || token != "" {
		users, err := h.Service.GetUsersByParams(username, token)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to retrieve user",
			})
		}
		if len(users) == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.JSON(fiber.Map{
			"count": len(users),
			"data":  users,
		})
	}

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
