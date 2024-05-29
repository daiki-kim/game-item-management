package controllers

import (
	"game-item-management/services"
)

type ITradeController interface {
}

type TradeController struct {
	service services.TradeService
}

func NewTradeController(service services.TradeService) ITradeController {
	return &TradeController{service: service}
}
