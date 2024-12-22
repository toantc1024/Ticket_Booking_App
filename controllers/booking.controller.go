package controllers

import (
	"tickets/models"
	"tickets/services"

	"github.com/gofiber/fiber/v2"
)

type BookingController struct {
	service services.BookingService
}

func NewBookingController(service services.BookingService) *BookingController {
	return &BookingController{
		service: service}
}

func (ctrl *BookingController) CreateBooking(c *fiber.Ctx) error {
	var bookingSchema models.CreateBookingSchema
	if err := c.BodyParser(&bookingSchema); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	booking, err := ctrl.service.CreateBooking(&bookingSchema)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(booking)
}

func (ctrl *BookingController) GetBookingByID(c *fiber.Ctx) error {
	bookingID := c.Params("id")

	booking, err := ctrl.service.GetBookingByID(bookingID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
	}

	return c.Status(fiber.StatusOK).JSON(booking)
}

func (ctrl *BookingController) UpdateBooking(c *fiber.Ctx) error {
	var booking models.Booking
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ctrl.service.UpdateBooking(&booking); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(booking)
}

func (ctrl *BookingController) DeleteBooking(c *fiber.Ctx) error {
	bookingID := c.Params("id")

	if err := ctrl.service.DeleteBooking(bookingID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
