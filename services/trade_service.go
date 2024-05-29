package services

import (
	"game-item-management/repositories"
)

type ITradeService interface {
}

type TradeService struct {
	repository repositories.ITradeRepository
}
