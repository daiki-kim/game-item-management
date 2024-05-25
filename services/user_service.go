package services

import (
	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	SignUp(inputUser dtos.CreateUserDTO) error
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository: repository}
}

func (s *UserService) SignUp(inputUser dtos.CreateUserDTO) error {
	hasshedPassword, err := bcrypt.GenerateFromPassword([]byte(inputUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Name:     inputUser.Name,
		Email:    inputUser.Email,
		Password: string(hasshedPassword),
	}

	err = s.repository.CreateNewUser(&newUser)
	if err != nil {
		return err
	}
	return nil
}
