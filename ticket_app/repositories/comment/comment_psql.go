package repositories

import (
	"gorm.io/gorm"
	"waizly.com/adriantoprasetyo/ticket_app/models"
	"waizly.com/adriantoprasetyo/ticket_app/repositories"
)

type psqlCommentRepo struct {
	Conn *gorm.DB
}

func NewPSQLComment(conn *gorm.DB) repositories.Comment {
	return &psqlCommentRepo{
		Conn: conn,
	}
}

func (p *psqlCommentRepo) Store(m models.Comment) (models.Comment, error) {

	err := p.Conn.Create(&m).Error
	if err != nil {
		return models.Comment{}, err
	}

	return m, nil
}

func (p *psqlCommentRepo) FetchByTicketId(id int) ([]models.Comment, error) {
	var (
		comment []models.Comment
		err     error
	)
	err = p.Conn.Where("ticket_id=?", id).Find(&comment).Error
	if err != nil {
		return []models.Comment{}, err
	}
	return comment, nil
}
