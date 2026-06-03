package service

import (
	"errors"
	"time"

	"sports-backend/Source/backend/internal/middleware"
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	_, err = s.repo.CreateUser(
		user.Username,
		string(hashedPassword),
	)

	return err
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(middleware.JwtSecret)
}

func (s *AuthService) UpdateUsername(userID int, username string) error {
	return s.repo.UpdateUsername(userID, username)
}

func (s *AuthService) DeleteUser(userID int) error {
	return s.repo.DeleteUser(userID)
}
