package services

import (
	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/repositories"
)

type IItemService interface {
	CreateNewItem(inputItem dtos.NewItemDTO, userId uint) (*models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) CreateNewItem(inputItem dtos.NewItemDTO, userId uint) (*models.Item, error) {
	newItem := models.Item{
		Name:        inputItem.Name,
		Description: inputItem.Description,
		UserID:      userId,
	}
	return s.repository.CreateNewItem(newItem)
}
