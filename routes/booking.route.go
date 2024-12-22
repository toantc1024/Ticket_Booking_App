package routes

import (
	"tickets/controllers"
	"tickets/repositories"
	"tickets/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BookingRoutes(app *fiber.App, db *gorm.DB) {
	bookingRepo := repositories.NewBookingRepository(db)
	bookingService := services.NewBookingService(bookingRepo)
	bookingController := controllers.NewBookingController(bookingService)

	bookingRoutes := app.Group("/bookings")
	bookingRoutes.Post("/create", bookingController.CreateBooking)
	bookingRoutes.Delete("/:id", bookingController.DeleteBooking)
}
