package repositories

import (
	"errors"
	"game-item-management/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	CreateNewItem(newItem models.Item) (*models.Item, error)
	FindAllItems() (*[]models.Item, error)
	FindItemById(itemId uint) (*models.Item, error)
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

func (r *ItemRepository) FindItemById(itemId uint) (*models.Item, error) {
	var foundItem models.Item
	result := r.db.First(&foundItem, itemId)
	if result.Error != nil {
		return nil, errors.New("item not found")
	}
	return &foundItem, nil
}
