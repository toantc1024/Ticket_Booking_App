package routes

import (
	"tickets/controllers"
	"tickets/repositories"
	"tickets/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TicketRoutes(app *fiber.App, db *gorm.DB) {
	ticketRepo := repositories.NewTicketRepository(db)
	ticketService := services.NewTicketService(ticketRepo)
	ticketController := controllers.NewTicketController(ticketService)

	ticketRoutes := app.Group("/tickets")
	ticketRoutes.Put("/:id", ticketController.UpdateTicket)
	ticketRoutes.Post("/create", ticketController.CreateTicket)
	ticketRoutes.Get("/all", ticketController.GetTickets)
}
