package repositories

import (
	"game-item-management/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateNewUser(newUser *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateNewUser(newUser *models.User) error {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var foundUser models.User
	result := r.db.First(&foundUser, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}
