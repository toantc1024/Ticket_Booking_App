package repositories

import (
	"gorm.io/gorm"
	"tickets/models"
)

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentDetails(paymentID int) (*models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetPaymentDetails(paymentID int) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, paymentID).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
