package routes

import (
	"tickets/controllers"
	"tickets/repositories"
	"tickets/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PaymentRoutes(app *fiber.App, db *gorm.DB) {
	paymentRepo := repositories.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService)

	paymentRoutes := app.Group("/payments")
	paymentRoutes.Post("/", paymentController.CreatePayment)
	paymentRoutes.Get("/:id", paymentController.GetPaymentDetails)
}
