package controllers

import (
	"game-item-management/dtos"
	"game-item-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	SignUp(c *gin.Context)
}

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) IUserController {
	return &UserController{service: service}
}

func (c *UserController) SignUp(ctx *gin.Context) {
	var newUserInput dtos.CreateUserDTO
	if err := ctx.ShouldBindJSON(&newUserInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.SignUp(newUserInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
