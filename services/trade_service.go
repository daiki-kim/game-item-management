package services

import (
	"errors"
	"fmt"
	"game-item-management/models"
	"game-item-management/repositories"
)

type ITradeService interface {
	CreateNewTrade(itemId, toUserId uint) (*models.Trade, error)
}

type TradeService struct {
	itemRepository  repositories.IItemRepository
	tradeRepository repositories.ITradeRepository
	userRepository  repositories.IUserRepository
	emailService    IEmailService
}

func NewTradeService(itemRepository repositories.IItemRepository, tradeRepository repositories.ITradeRepository) ITradeService {
	return &TradeService{itemRepository: itemRepository, tradeRepository: tradeRepository}
}

func (s *TradeService) CreateNewTrade(itemId, toUserId uint) (*models.Trade, error) {
	item, err := s.itemRepository.FindItemById(itemId)
	if err != nil {
		return nil, errors.New("item not found")
	}
	trade := models.Trade{
		Is_Accepted: false,
		ItemID:      itemId,
		FromUserID:  item.UserID,
		ToUserID:    toUserId,
	}
	newTrade, _ := s.tradeRepository.CreateNewTrade(trade)

	// get user models and send email
	toUser, _ := s.userRepository.FindById(trade.ToUserID)
	fromUser, _ := s.userRepository.FindById(trade.FromUserID)
	subject := fmt.Sprintf("Trade request from %s", toUser.Name)
	body := fmt.Sprintf("Trade request from %s. Please accept or decline.", toUser.Name)
	s.emailService.SendEmail(fromUser.Email, subject, body)

	return newTrade, nil
}
