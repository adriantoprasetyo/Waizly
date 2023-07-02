package models

import "time"

func (Ticket) TableName() string {
	return "tickets"
}

type Ticket struct {
	Id        int       `json:"id" gorm:"column:id"`
	UserId    int       `json:"user_id" gorm:"column:user_id"`
	AgentId   int       `json:"agent_id" gorm:"column:agent_id"`
	Status    bool      `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type HandoverTicket struct {
	TicketId int `json:"ticket_id" gorm:"column:ticket_id"`
	AgentId  int `json:"agent_id" gorm:"column:agent_id"`
}
