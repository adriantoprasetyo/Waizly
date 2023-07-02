package controllers

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"waizly.com/adriantoprasetyo/ticket_app/models"
	repoComment "waizly.com/adriantoprasetyo/ticket_app/repositories/comment"
	repoTicket "waizly.com/adriantoprasetyo/ticket_app/repositories/ticket"
)

type CommentController struct {
	DB *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{
		DB: db,
	}
}

func (tc *CommentController) Store(ctx *gin.Context) {
	var (
		req     models.StoreComment
		ticket  models.Ticket
		comment models.Comment
		err     error
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketId, err := strconv.Atoi(ctx.Query("ticket_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketRepo := repoTicket.NewPSQLTicket(tc.DB)
	ticket, err = ticketRepo.Get(ticketId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment = models.Comment{}
	comment.Text = req.Text
	comment.TicketId = ticket.Id
	comment.UserId = ticket.UserId

	commentRepo := repoComment.NewPSQLComment(tc.DB)
	comment, err = commentRepo.Store(comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := make(map[string]interface{})
	result["ticket_id"] = ticket.Id
	result["comment"] = comment

	ctx.JSON(http.StatusOK, result)
}
