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
	GetUserFromToken(tokenString string) (*models.User, error)
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

	return CreateToken(foundUser.ID, foundUser.Email)
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
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
	if err != nil {
		return nil, err
	}

	var gotUser *models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}

		gotUser, err = s.repository.FindByEmail(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return gotUser, nil
}
