package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"waizly.com/adriantoprasetyo/ticket_app/controllers"
	"waizly.com/adriantoprasetyo/ticket_app/utility"
)

func main() {
	r := gin.Default()

	db := utility.ConnectDataBase()

	userCtrl := controllers.NewUserController(db)
	ticketCtrl := controllers.NewTicketController(db)
	commentCtrl := controllers.NewCommentController(db)

	auth := r.Group("/auth")
	{
		auth.POST("/register", userCtrl.Register)
		auth.POST("/login", userCtrl.Login)
		auth.DELETE("/user/delete/:id", userCtrl.Delete)
	}

	publicTicket := r.Group("/ticket")
	{
		publicTicket.POST("", ticketCtrl.Store)
		publicTicket.GET("/:id", ticketCtrl.Get)
	}

	protectedTicket := r.Group("/ticket")
	protectedTicket.Use(JwtAuthMiddleware())
	{
		protectedTicket.GET("", ticketCtrl.Fetch)
		protectedTicket.PUT("/handover", ticketCtrl.Handover)
		protectedTicket.PUT("/close/:id", ticketCtrl.CloseTicket)
	}

	protectedComment := r.Group("/comment")
	protectedComment.Use(JwtAuthMiddleware())
	{
		protectedComment.POST("/:ticket_id", commentCtrl.Store)
	}

	r.Run(":8080")
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utility.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
