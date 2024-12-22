package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	gorm.Model
	TicketID  string  `gorm:"type:char(36);primary_key" json:"ticket_id,omitempty"`
	Title     string  `gorm:"type:varchar(255);not null" json:"title"`
	Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	SeatPlace string  `gorm:"type:varchar(255);not null" json:"seatPlace"`
}

func (ticket *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	ticket.TicketID = uuid.New().String()
	ticket.CreatedAt = time.Now()
	return nil
}

type CreateTicketSchema struct {
	Title     string  `json:"title" validate:"required,min=3,max=255"`
	Price     float64 `json:"price" validate:"required,gt=0"`
	SeatPlace string  `json:"seatPlace" validate:"required,min=1,max=255"`
}

type UpdateTicketSchema struct {
	Title     string  `json:"title" validate:"omitempty,min=3,max=255"`
	Price     float64 `json:"price" validate:"omitempty,gt=0"`
	SeatPlace string  `json:"seatPlace" validate:"omitempty,min=1,max=255"`
}
