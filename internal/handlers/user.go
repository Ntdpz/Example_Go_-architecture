package handlers

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/internal/services"
	"Example_Go_architecture/models"
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

func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	token := c.Query("Token")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token is required",
		})
	}

	var userData models.Users

	// ดึงข้อมูลจาก Body (JSON)
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// อัปเดตข้อมูลผู้ใช้
	updatedUser, err := h.Service.UpdateUserByToken(token, userData)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"data":    updatedUser,
		"message": "User updated successfully",
	})
}

func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	token := c.Query("Token")

	if token == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Token is required",
		})
	}
	err := h.Service.DeleteUserByToken(token)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
