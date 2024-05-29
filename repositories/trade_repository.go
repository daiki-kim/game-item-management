package repositories

import (
	"game-item-management/models"

	"gorm.io/gorm"
)

type ITradeRepository interface {
	CreateNewTrade(newTrade models.Trade) (*models.Trade, error)
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
