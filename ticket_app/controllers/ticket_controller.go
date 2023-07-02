package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"waizly.com/adriantoprasetyo/ticket_app/models"
	repoComment "waizly.com/adriantoprasetyo/ticket_app/repositories/comment"
	repoTicket "waizly.com/adriantoprasetyo/ticket_app/repositories/ticket"
	repoUser "waizly.com/adriantoprasetyo/ticket_app/repositories/user"
	"waizly.com/adriantoprasetyo/ticket_app/utility"
)

type TicketController struct {
	DB *gorm.DB
}

func NewTicketController(db *gorm.DB) *TicketController {
	return &TicketController{
		DB: db,
	}
}

func (tc *TicketController) Fetch(ctx *gin.Context) {

	user_id, err := utility.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepo := repoUser.NewPSQLUser(tc.DB)
	user, err := userRepo.Get(int(user_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := make(map[string]interface{})
	if user.Role == "supervisor" {
		params["agent_id"] = ctx.Query("agent_id")
		params["ticket_status"] = ctx.DefaultQuery("ticket_status", "true")
	} else {
		params["agent_id"] = user_id
	}

	ticketRepo := repoTicket.NewPSQLTicket(tc.DB)
	data, err := ticketRepo.Fetch(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, data)
}

func (tc *TicketController) Store(ctx *gin.Context) {
	var (
		req     models.StoreComment
		user    models.User
		ticket  models.Ticket
		comment models.Comment
		err     error
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = models.User{}
	user.Username = fmt.Sprintf("%v", uuid.NewV4())

	userRepo := repoUser.NewPSQLUser(tc.DB)
	user, err = userRepo.Register(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket = models.Ticket{}
	ticket.UserId = user.Id

	ticketRepo := repoTicket.NewPSQLTicket(tc.DB)
	ticket, err = ticketRepo.Store(ticket)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment = models.Comment{}
	comment.Text = req.Text
	comment.TicketId = ticket.Id
	comment.UserId = user.Id

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

func (tc *TicketController) Get(ctx *gin.Context) {
	var (
		ticket models.Ticket
		err    error
	)

	ticketId, err := strconv.Atoi(ctx.Query("id"))
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

	if !ticket.Status {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
	}

	ctx.JSON(http.StatusOK, ticket)
}

func (tc *TicketController) Handover(ctx *gin.Context) {
	var (
		req    models.HandoverTicket
		ticket models.Ticket
		err    error
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, err := utility.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepo := repoUser.NewPSQLUser(tc.DB)
	user, err := userRepo.Get(int(user_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Role != "supervisor" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "you do not have access to this API"})
		return
	}

	ticketRepo := repoTicket.NewPSQLTicket(tc.DB)
	ticket, err = ticketRepo.Handover(models.Ticket{Id: req.TicketId}, req.AgentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ticket)

}

func (tc *TicketController) CloseTicket(ctx *gin.Context) {
	var (
		ticket models.Ticket
		err    error
	)

	ticketId, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketRepo := repoTicket.NewPSQLTicket(tc.DB)
	err = ticketRepo.Close(models.Ticket{Id: ticketId}, models.Ticket{Status: false})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ticket.Status {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
	}

	ctx.JSON(http.StatusOK, ticket)
}
