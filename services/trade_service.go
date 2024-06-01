package services

import (
	"errors"
	"fmt"
	"log"

	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/repositories"
)

type ITradeService interface {
	CreateNewTrade(itemId, toUserId uint) (*models.Trade, error)
	FindTradeByTradeId(tradeId uint) (*models.Trade, error)
	UpdateTradeStatus(tradeId, userId uint, inputTrade dtos.UpdateTradeDTO) (*models.Trade, error)
}

type TradeService struct {
	itemRepository  repositories.IItemRepository
	tradeRepository repositories.ITradeRepository
	userRepository  repositories.IUserRepository
	emailService    IEmailService
}

func NewTradeService(
	itemRepository repositories.IItemRepository,
	tradeRepository repositories.ITradeRepository,
	userRepository repositories.IUserRepository,
	emailService IEmailService,
) ITradeService {
	return &TradeService{
		itemRepository:  itemRepository,
		tradeRepository: tradeRepository,
		userRepository:  userRepository,
		emailService:    emailService,
	}
}

func (s *TradeService) CreateNewTrade(itemId, toUserId uint) (*models.Trade, error) {
	item, err := s.itemRepository.FindItemById(itemId)
	if err != nil {
		return nil, errors.New("item not found")
	}
	trade := models.Trade{
		ItemID:     itemId,
		FromUserID: item.UserID,
		ToUserID:   toUserId,
	}
	newTrade, _ := s.tradeRepository.CreateNewTrade(trade)

	// get user models and send email
	toUser, _ := s.userRepository.FindById(trade.ToUserID)
	fromUser, _ := s.userRepository.FindById(trade.FromUserID)
	subject := fmt.Sprintf("Trade request from %s", toUser.Name)
	body := fmt.Sprintf("Trade request from %s. Please accept or decline.", toUser.Name)
	s.emailService.SendEmail(fromUser.Email, subject, body)
	log.Printf("email sent to %s.\nsubject: %s\nbody: %s", fromUser.Email, subject, body)

	return newTrade, nil
}

func (s *TradeService) FindTradeByTradeId(tradeId uint) (*models.Trade, error) {
	return s.tradeRepository.FindTradeByTradeId(tradeId)
}

func (s *TradeService) UpdateTradeStatus(tradeId, userId uint, inputTrade dtos.UpdateTradeDTO) (*models.Trade, error) {
	panic("")
}
