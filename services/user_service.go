package services

import (
	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Signup(inputUser dtos.SignupUserDTO) error
	Login(email string, password string) (*string, error)
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository: repository}
}

func (s *UserService) Signup(inputUser dtos.SignupUserDTO) error {
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

func (s *UserService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &foundUser.Email, nil // TODO: Change to token after creating Token generate func
}
