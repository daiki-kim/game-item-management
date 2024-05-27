package controllers

import (
	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	CreateItem(ctx *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (c *ItemController) CreateItem(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	modelsUser, ok := user.(*models.User)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId := modelsUser.ID

	var inputItem dtos.NewItemDTO
	if err := ctx.ShouldBindJSON(&inputItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdItem, err := c.service.CreateNewItem(inputItem, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}
	ctx.JSON(http.StatusCreated, createdItem)
}
