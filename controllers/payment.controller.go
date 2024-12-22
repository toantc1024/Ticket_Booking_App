package controllers

import (
	"tickets/models"
	"tickets/services"

	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	service services.PaymentService
}

func NewPaymentController(service services.PaymentService) *PaymentController {
	return &PaymentController{service}
}

func (ctrl *PaymentController) CreatePayment(c *fiber.Ctx) error {
	var payment models.Payment
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ctrl.service.CreatePayment(&payment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(payment)
}

func (ctrl *PaymentController) GetPaymentDetails(c *fiber.Ctx) error {
	paymentID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payment ID"})
	}

	payment, err := ctrl.service.GetPaymentDetails(paymentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Payment not found"})
	}

	return c.Status(fiber.StatusOK).JSON(payment)
}
