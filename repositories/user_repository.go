package repositories

import (
	"game-item-management/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateNewUser(newUser *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateNewUser(newUser *models.User) error {
	return r.db.Create(&newUser).Error
}
