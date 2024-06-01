package repositories

import (
	"errors"
	"game-item-management/models"

	"gorm.io/gorm"
)

type ITradeRepository interface {
	CreateNewTrade(newTrade models.Trade) (*models.Trade, error)
	FindTradeByTradeId(tradeId uint) (*models.Trade, error)
	UpdateTrade(trade models.Trade) (*models.Trade, error)
	FindAllTradesByItemId(itemId uint) (*[]models.Trade, error)
}

type TradeRepository struct {
	db *gorm.DB
}

func NewTradeRepository(db *gorm.DB) ITradeRepository {
	return &TradeRepository{db: db}
}

func (r *TradeRepository) CreateNewTrade(newTrade models.Trade) (*models.Trade, error) {
	result := r.db.Create(&newTrade)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newTrade, nil
}

func (r *TradeRepository) FindTradeByTradeId(tradeId uint) (*models.Trade, error) {
	var foundTrade models.Trade
	result := r.db.First(&foundTrade, tradeId)
	if result.Error != nil {
		return nil, errors.New("trade not found")
	}
	return &foundTrade, nil
}

func (r *TradeRepository) UpdateTrade(trade models.Trade) (*models.Trade, error) {
	result := r.db.Save(&trade)
	if result.Error != nil {
		return nil, result.Error
	}
	return &trade, nil
}

func (r *TradeRepository) FindAllTradesByItemId(itemId uint) (*[]models.Trade, error) {
	var foundTrades []models.Trade
	result := r.db.Where("item_id = ?", itemId).Find(&foundTrades)
	if result.Error != nil {
		return nil, errors.New("trades not found")
	}
	return &foundTrades, nil
}
