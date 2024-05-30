package services

import (
	"game-item-management/models"
	"game-item-management/repositories"
)

type ITradeService interface {
	CreateNewTrade(itemId, toUserId uint) (*models.Trade, error)
}

type TradeService struct {
	itemRepository  repositories.IItemRepository
	tradeRepository repositories.ITradeRepository
}

func NewTradeService(itemRepository repositories.IItemRepository, tradeRepository repositories.ITradeRepository) ITradeService {
	return &TradeService{itemRepository: itemRepository, tradeRepository: tradeRepository}
}

func (s *TradeService) CreateNewTrade(itemId, toUserId uint) (*models.Trade, error) {
	item, err := s.itemRepository.FindItemById(itemId)
	if err != nil {
		return nil, err
	}

	trade := models.Trade{
		Is_Accepted: false,
		ItemID:      itemId,
		FromUserID:  item.UserID,
		ToUserID:    toUserId,
	}
	return s.tradeRepository.CreateNewTrade(trade)
}
