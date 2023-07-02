package repositories

import (
	"waizly.com/adriantoprasetyo/ticket_app/models"
)

type User interface {
	Register(m models.User) (models.User, error)
	LoginCheck(username string, password string) (string, error)
	GetByUsername(username string) (int64, error)
	Get(id int) (models.User, error)
	Delete(id int) error
}
