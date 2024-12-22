package services

import (
	"tickets/models"
	"tickets/repositories"
)

type TicketService interface {
	CreateTicket(ticketSchema *models.CreateTicketSchema) (*models.Ticket, error)
	GetTicketByID(ticketID string) (*models.Ticket, error)
	GetTickets() ([]models.Ticket, error)
	UpdateTicket(ticket *models.Ticket) error
	DeleteTicket(ticketID string) error
}

type ticketService struct {
	repo repositories.TicketRepository
}

func NewTicketService(repo repositories.TicketRepository) TicketService {
	return &ticketService{repo}
}

func (s *ticketService) GetTickets() ([]models.Ticket, error) {
	return s.repo.GetTickets()
}

func (s *ticketService) CreateTicket(ticketSchema *models.CreateTicketSchema) (*models.Ticket, error) {
	ticket := &models.Ticket{
		Title:     ticketSchema.Title,
		Price:     ticketSchema.Price,
		SeatPlace: ticketSchema.SeatPlace,
	}
	err := s.repo.CreateTicket(ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *ticketService) GetTicketByID(ticketID string) (*models.Ticket, error) {
	return s.repo.GetTicketByID(ticketID)
}

func (s *ticketService) UpdateTicket(ticket *models.Ticket) error {
	return s.repo.UpdateTicket(ticket)
}

func (s *ticketService) DeleteTicket(ticketID string) error {
	return s.repo.DeleteTicket(ticketID)
}
