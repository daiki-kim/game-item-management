package services

import (
	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/repositories"
)

type IItemService interface {
	CreateNewItem(inputItem dtos.NewItemDTO, userId uint) (*models.Item, error)
	FindAllItems() (*[]models.Item, error)
	FindItemById(itemId uint) (*models.Item, error)
	UpdateItem(itemId uint, inputItem dtos.UpdateItemDTO) (*models.Item, error)
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

func (s *ItemService) FindAllItems() (*[]models.Item, error) {
	return s.repository.FindAllItems()
}

func (s *ItemService) FindItemById(itemId uint) (*models.Item, error) {
	return s.repository.FindItemById(itemId)
}

func (s *ItemService) UpdateItem(itemId uint, inputItem dtos.UpdateItemDTO) (*models.Item, error) {
	targetItem, err := s.repository.FindItemById(itemId)
	if err != nil {
		return nil, err
	}
	if inputItem.Name != nil {
		targetItem.Name = *inputItem.Name
	}
	if inputItem.Description != nil {
		targetItem.Description = *inputItem.Description
	}
	return s.repository.UpdateItem(*targetItem)
}
