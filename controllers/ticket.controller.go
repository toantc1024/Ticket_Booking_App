package controllers

import (
	"tickets/models"
	"tickets/services"

	"github.com/gofiber/fiber/v2"
)

type TicketController struct {
	service services.TicketService
}

func NewTicketController(service services.TicketService) *TicketController {
	return &TicketController{
		service: service}
}

func (ctrl *TicketController) CreateTicket(c *fiber.Ctx) error {
	var ticketSchema models.CreateTicketSchema
	if err := c.BodyParser(&ticketSchema); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ticket, err := ctrl.service.CreateTicket(&ticketSchema)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(ticket)
}

func (ctrl *TicketController) GetTickets(c *fiber.Ctx) error {
	tickets, err := ctrl.service.GetTickets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(tickets)
}

func (ctrl *TicketController) GetTicketByID(c *fiber.Ctx) error {
	ticketID := c.Params("id")

	ticket, err := ctrl.service.GetTicketByID(ticketID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ticket not found"})
	}

	return c.Status(fiber.StatusOK).JSON(ticket)
}

func (ctrl *TicketController) UpdateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket
	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ctrl.service.UpdateTicket(&ticket); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(ticket)
}

func (ctrl *TicketController) DeleteTicket(c *fiber.Ctx) error {
	ticketID := c.Params("id")

	if err := ctrl.service.DeleteTicket(ticketID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
