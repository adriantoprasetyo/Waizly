package repositories

import (
	"waizly.com/adriantoprasetyo/ticket_app/models"
)

type Comment interface {
	Store(m models.Comment) (models.Comment, error)
	FetchByTicketId(id int) ([]models.Comment, error)
}
