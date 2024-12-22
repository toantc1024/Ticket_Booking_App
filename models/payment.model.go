package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	PaymentID string  `gorm:"type:char(36);primary_key" json:"payment_id,omitempty"`
	Amount    float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
}

func (payment *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	payment.PaymentID = uuid.New().String()
	payment.CreatedAt = time.Now()
	return nil
}

type CreatePaymentSchema struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type UpdatePaymentSchema struct {
	Amount float64 `json:"amount" validate:"omitempty,gt=0"`
}

type PaymentResponseSchema struct {
	PaymentID string    `json:"payment_id,omitempty"`
	Amount    float64   `json:"amount,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
