package controllers

import (
	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITradeController interface {
	CreateNewTrade(ctx *gin.Context)
	FindTradeByTradeId(ctx *gin.Context)
	UpdateTradeStatus(ctx *gin.Context)
	FindAllTradesByItemId(ctx *gin.Context)
	FindAllTradesByUserId(ctx *gin.Context)
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

func (c *TradeController) UpdateTradeStatus(ctx *gin.Context) {
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

	var inputTrade dtos.UpdateTradeDTO
	if err := ctx.ShouldBindJSON(&inputTrade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tradeId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid trade id"})
		return
	}
	updateTrade, err := c.service.UpdateTradeStatus(uint(tradeId), userId, inputTrade)
	if err != nil {
		switch err.Error() {
		case "trade not found":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "you are not the owner of this trade":
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		}
		return
	}
	ctx.JSON(http.StatusOK, updateTrade)
}

func (c *TradeController) FindAllTradesByItemId(ctx *gin.Context) {
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

	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}
	foundTrades, err := c.service.FindAllTradesByItemId(uint(itemId))
	if err != nil {
		if err.Error() == "trades not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, foundTrades)
}

func (c *TradeController) FindAllTradesByUserId(ctx *gin.Context) {
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

	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	foundTrades, err := c.service.FindAllTradesByUserId(uint(userId))
	if err != nil {
		if err.Error() == "trades not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, foundTrades)
}
