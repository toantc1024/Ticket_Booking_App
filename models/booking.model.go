package models

import (
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type StatusEnum string

const (
	Pending   StatusEnum = "pending"
	Confirmed StatusEnum = "confirmed"
	Canceled  StatusEnum = "canceled"
)

type Booking struct {
	gorm.Model
	BookingID   string     `gorm:"type:char(36);primary_key" json:"booking_id,omitempty"`
	BookingDate time.Time  `gorm:"not null" json:"bookingDate"`
	Status      StatusEnum `gorm:"type:varchar(20);not null" json:"status"`
	UserID      string     `gorm:"not null" json:"userId"`
	User        User       `gorm:"foreignKey:UserID" json:"user"`
	TicketID    string     `gorm:"not null" json:"ticketId"`
	Ticket      Ticket     `gorm:"foreignKey:TicketID" json:"ticket"`
	PaymentID   string     `gorm:"not null" json:"paymentId"`
	Payment     Payment    `gorm:"foreignKey:PaymentID" json:"payment"`
}

func (booking *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	booking.BookingID = uuid.New().String()
	booking.CreatedAt = time.Now()
	return nil
}

type BookingResponseSchema struct {
	BookingID   string     `json:"booking_id,omitempty"`
	BookingDate time.Time  `json:"bookingDate,omitempty"`
	Status      StatusEnum `json:"status,omitempty"`
	UserID      string     `json:"user_id,omitempty"`
	User        User       `json:"user,omitempty"`
	TicketID    string     `json:"ticket_id,omitempty"`
	Ticket      Ticket     `json:"ticket,omitempty"`
	PaymentID   string     `json:"payment_id,omitempty"`
	Payment     Payment    `json:"payment,omitempty"`
}

type CreateBookingSchema struct {
	BookingDate time.Time  `json:"bookingDate" validate:"required"`
	Status      StatusEnum `json:"status" validate:"required"`
	UserID      string     `json:"user_id" validate:"required"`
	TicketID    string     `json:"ticket_id" validate:"required"`
	PaymentID   string     `json:"payment_id"`
}

type BookingFilterSchema struct {
	Status StatusEnum `json:"status,omitempty"`
}

type BookingUpdateSchema struct {
	Status StatusEnum `json:"status,omitempty"`
}
