package services

import (
	"tickets/models"
	"tickets/repositories"
)

type BookingService interface {
	CreateBooking(bookingSchema *models.CreateBookingSchema) (*models.Booking, error)
	GetBookingByID(bookingID string) (*models.Booking, error)
	UpdateBooking(booking *models.Booking) error
	DeleteBooking(bookingID string) error
}

type bookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) BookingService {
	return &bookingService{repo}
}

func (s *bookingService) CreateBooking(bookingSchema *models.CreateBookingSchema) (*models.Booking, error) {
	booking := &models.Booking{
		BookingDate: bookingSchema.BookingDate,
		Status:      bookingSchema.Status,
		UserID:      bookingSchema.UserID,
		TicketID:    bookingSchema.TicketID,
		PaymentID:   bookingSchema.PaymentID,
	}
	err := s.repo.CreateBooking(booking)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) GetBookingByID(bookingID string) (*models.Booking, error) {
	return s.repo.GetBookingByID(bookingID)
}

func (s *bookingService) UpdateBooking(booking *models.Booking) error {
	return s.repo.UpdateBooking(booking)
}

func (s *bookingService) DeleteBooking(bookingID string) error {
	return s.repo.DeleteBooking(bookingID)
}
