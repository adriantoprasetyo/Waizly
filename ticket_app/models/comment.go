package models

import "time"

func (Comment) TableName() string {
	return "comments"
}

type Comment struct {
	Id        int       `json:"id" gorm:"column:id"`
	TicketId  int       `json:"ticket_id" gorm:"column:ticket_id"`
	UserId    int       `json:"user_id" gorm:"column:user_id"`
	Text      string    `json:"text" gorm:"column:text"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type StoreComment struct {
	Text string `json:"text" gorm:"column:text" binding:"required"`
}
