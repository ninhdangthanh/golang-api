package services

import (
	"log"
	"net/http"

	"github.com/example/intern/database"
	"github.com/example/intern/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{db: database.GetDB()}
}

func (s *UserService) CreateUser(user *models.UserModel) error {
	var existingUser models.UserModel
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return &AppError{StatusCode: http.StatusBadRequest, Message: "Email is already taken"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return &AppError{StatusCode: http.StatusInternalServerError, Message: "Failed to process password"}
	}
	user.Password = string(hashedPassword)

	if err := s.db.Create(user).Error; err != nil {
		return &AppError{StatusCode: http.StatusInternalServerError, Message: "Failed to create user"}
	}

	return nil
}

func (s *UserService) AuthenticateUser(email, password string) (*models.UserModel, error) {
	var user models.UserModel

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}

// AppError represents an application error with a status code and message
type AppError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}
