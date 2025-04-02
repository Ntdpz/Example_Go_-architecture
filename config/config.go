package config

import (
	"Example_Go_architecture/database"
	"Example_Go_architecture/internal/handlers"
	"Example_Go_architecture/internal/repositories"
	"Example_Go_architecture/internal/services"
)

func InitializeServices() *handlers.UserHandler {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	return userHandler
}
