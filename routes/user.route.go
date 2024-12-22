package routes

import (
	"tickets/controllers"
	"tickets/repositories"
	"tickets/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userRoutes := app.Group("/users")
	userRoutes.Post("/register", userController.CreateUser)
	userRoutes.Post("/login", userController.Login)
	userRoutes.Post("/logout/:id", userController.Logout)
}
