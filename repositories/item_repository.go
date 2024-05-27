package repositories

import (
	"game-item-management/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	CreateNewItem(newItem models.Item) (*models.Item, error)
	FindAllItems() (*[]models.Item, error)
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

func (r *ItemRepository) FindAllItems() (*[]models.Item, error) {
	var foundItems []models.Item
	result := r.db.Find(&foundItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundItems, nil
}
