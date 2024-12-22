package repositories

import (
	"gorm.io/gorm"
	"tickets/models"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking) error
	GetBookingByID(bookingID string) (*models.Booking, error)
	UpdateBooking(booking *models.Booking) error
	DeleteBooking(bookingID string) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingByID(bookingID string) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.Preload("User").Preload("Ticket").Preload("Payment").First(&booking, "booking_id = ?", bookingID).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) UpdateBooking(booking *models.Booking) error {
	return r.db.Save(booking).Error
}

func (r *bookingRepository) DeleteBooking(bookingID string) error {
	return r.db.Delete(&models.Booking{}, "booking_id = ?", bookingID).Error
}
