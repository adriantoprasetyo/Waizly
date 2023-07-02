package controllers

import (
	"html"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"waizly.com/adriantoprasetyo/ticket_app/models"
	repo "waizly.com/adriantoprasetyo/ticket_app/repositories/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var (
		req  models.RegisterInput
		user models.User
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataRepo := repo.NewPSQLUser(uc.DB)
	count, _ := dataRepo.GetByUsername(req.Username)
	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already being used"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user = models.User{}

	user.Password = string(hashedPassword)

	//remove spaces in username
	user.Username = html.EscapeString(strings.TrimSpace(req.Username))
	user.Role = req.Role

	_, err = dataRepo.Register(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var input models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	dataRepo := repo.NewPSQLUser(uc.DB)
	token, err := dataRepo.LoginCheck(u.Username, u.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) Delete(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataRepo := repo.NewPSQLUser(uc.DB)
	err = dataRepo.Delete(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete user success"})
}
