package services

import (
	"tickets/models"
	"tickets/repositories"
)

type PaymentService interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentDetails(paymentID int) (*models.Payment, error)
}

type paymentService struct {
	repo repositories.PaymentRepository
}

func NewPaymentService(repo repositories.PaymentRepository) PaymentService {
	return &paymentService{repo}
}

func (s *paymentService) CreatePayment(payment *models.Payment) error {
	return s.repo.CreatePayment(payment)
}

func (s *paymentService) GetPaymentDetails(paymentID int) (*models.Payment, error) {
	return s.repo.GetPaymentDetails(paymentID)
}
