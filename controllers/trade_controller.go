package controllers

import (
	"game-item-management/models"
	"game-item-management/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITradeController interface {
	CreateNewTrade(ctx *gin.Context)
	FindTradeByTradeId(ctx *gin.Context)
}

type TradeController struct {
	service services.ITradeService
}

func NewTradeController(service services.ITradeService) ITradeController {
	return &TradeController{service: service}
}

func (c *TradeController) CreateNewTrade(ctx *gin.Context) {
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
	toUserId := modelsUser.ID

	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}
	newItemTrade, err := c.service.CreateNewTrade(uint(itemId), toUserId)
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}
	ctx.JSON(http.StatusCreated, newItemTrade)
}

func (c *TradeController) FindTradeByTradeId(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	_, ok := user.(*models.User)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tradeId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid trade id"})
		return
	}
	foundTrade, err := c.service.FindTradeByTradeId(uint(tradeId))
	if err != nil {
		if err.Error() == "trade not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, foundTrade)
}
