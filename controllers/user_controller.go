package controllers

import (
	"game-item-management/dtos"
	"game-item-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) IUserController {
	return &UserController{service: service}
}

func (c *UserController) Signup(ctx *gin.Context) {
	var newUserInput dtos.SignupUserDTO
	if err := ctx.ShouldBindJSON(&newUserInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Signup(newUserInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *UserController) Login(ctx *gin.Context) {
	var userInput dtos.LoginUserDTO
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userToken, err := c.service.Login(userInput.Email, userInput.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": userToken})
}
