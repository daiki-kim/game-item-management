package repositories

import (
	"game-item-management/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	CreateNewItem(newItem models.Item) (*models.Item, error)
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) CreateNewItem(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}
