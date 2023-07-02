package repositories

import (
	"waizly.com/adriantoprasetyo/ticket_app/models"
)

type Ticket interface {
	Fetch(filter map[string]interface{}) ([]models.Ticket, error)
	Store(m models.Ticket) (models.Ticket, error)
	Get(id int) (models.Ticket, error)
	Handover(m models.Ticket, data interface{}) (models.Ticket, error)
	Close(m models.Ticket, data interface{}) error
}
