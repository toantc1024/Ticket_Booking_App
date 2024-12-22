package repositories

import (
	"gorm.io/gorm"
	"tickets/models"
)

type TicketRepository interface {
	CreateTicket(ticket *models.Ticket) error
	GetTicketByID(ticketID string) (*models.Ticket, error)
	UpdateTicket(ticket *models.Ticket) error
	GetTickets() ([]models.Ticket, error)
	DeleteTicket(ticketID string) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) CreateTicket(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) GetTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(ticketID string) (*models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.First(&ticket, "ticket_id = ?", ticketID).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) UpdateTicket(ticket *models.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) DeleteTicket(ticketID string) error {
	return r.db.Delete(&models.Ticket{}, "ticket_id = ?", ticketID).Error
}
