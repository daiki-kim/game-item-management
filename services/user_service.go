package services

import (
	"fmt"
	"game-item-management/dtos"
	"game-item-management/models"
	"game-item-management/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Signup(inputUser dtos.SignupUserDTO) error
	Login(email string, password string) (*string, error)
	GetUsersProfile(name string) (*[]models.User, error)
	GetUserFromToken(tokenString string) (*models.User, error)
	FindUserById(userId uint) (*models.User, error)
	UpdateUserProfile(updateUserId uint, inputUser dtos.UpdateUserDTO) error
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
		Name:        inputUser.Name,
		Email:       inputUser.Email,
		Password:    string(hasshedPassword),
		Description: inputUser.Description,
	}

	err = s.repository.CreateNewUser(newUser)
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

	return CreateToken(foundUser.ID, foundUser.Email)
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *UserService) GetUserFromToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		fmt.Printf("Invalid token: %v\n", err)
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	var gotUser *models.User
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}

		gotUser, err = s.repository.FindByEmail(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	fmt.Printf("Token: %s\n", tokenString)
	fmt.Printf("Parsed Token: %+v\n", token)
	fmt.Printf("Claims: %+v\n", claims)

	return gotUser, nil
}

func (s *UserService) GetUsersProfile(name string) (*[]models.User, error) {
	foundUsers, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return foundUsers, nil
}

func (s *UserService) FindUserById(userId uint) (*models.User, error) {
	return s.repository.FindById(userId)
}

func (s *UserService) UpdateUserProfile(updateUserId uint, inputUser dtos.UpdateUserDTO) error {
	targetUser, err := s.FindUserById(updateUserId)
	if err != nil {
		return err
	}
	if inputUser.Name != nil {
		targetUser.Name = *inputUser.Name
	}
	if inputUser.Email != nil {
		targetUser.Email = *inputUser.Email
	}
	if inputUser.Description != nil {
		targetUser.Description = *inputUser.Description
	}
	return s.repository.Update(*targetUser)
}
