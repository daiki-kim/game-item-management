package repositories

import (
	"gorm.io/gorm"
)

type ITradeRepository interface {
}

type TradeRepository struct {
	db *gorm.DB
}

func NewTradeRepository(db *gorm.DB) ITradeRepository {
	return &TradeRepository{db: db}
}
