package repositories

import (
	"gorm.io/gorm"
	"waizly.com/adriantoprasetyo/ticket_app/models"
	"waizly.com/adriantoprasetyo/ticket_app/repositories"
)

type psqlTicketRepo struct {
	Conn *gorm.DB
}

func NewPSQLTicket(conn *gorm.DB) repositories.Ticket {
	return &psqlTicketRepo{
		Conn: conn,
	}
}

func (p *psqlTicketRepo) Fetch(filter map[string]interface{}) ([]models.Ticket, error) {

	var (
		err  error
		list []models.Ticket
	)

	db := p.Conn

	if v, ok := filter["agent_id"]; ok {
		db = db.Where("agent_id=?", v)
	}

	if v, ok := filter["ticket_status"]; ok {
		db = db.Where("status=?", v)
	}

	if err = db.Find(&list).Error; err != nil {
		return list, err
	}

	return list, nil
}

func (p *psqlTicketRepo) Store(m models.Ticket) (models.Ticket, error) {

	m.Status = true
	err := p.Conn.Create(&m).Error
	if err != nil {
		return models.Ticket{}, err
	}

	return m, nil
}

func (p *psqlTicketRepo) Get(id int) (models.Ticket, error) {
	var (
		ticket models.Ticket
		err    error
	)
	err = p.Conn.Where("id=?", id).Find(&ticket).Error
	if err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (p *psqlTicketRepo) Handover(m models.Ticket, data interface{}) (models.Ticket, error) {

	res := p.Conn.Model(&m).Where(m).Updates(data)
	if res.Error != nil {
		return m, res.Error
	}

	return models.Ticket{}, nil
}

func (p *psqlTicketRepo) Close(m models.Ticket, data interface{}) error {
	res := p.Conn.Model(&m).Where(m).Updates(data)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
