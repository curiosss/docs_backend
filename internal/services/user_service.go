package services

import (
	"errors"
	"time"

	"docs-notify/internal/dto"
	"docs-notify/internal/models"
	"docs-notify/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *dto.RegisterUserRequest) (*dto.UserResponse, error)
	Login(req *dto.LoginUserRequest) (*dto.LoginResponse, error)
}

type userService struct {
	repo      repositories.UserRepository
	jwtSecret string
}

func NewUserService(repo repositories.UserRepository, jwtSecret string) UserService {
	return &userService{repo, jwtSecret}
}

func (s *userService) Register(req *dto.RegisterUserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{ID: user.ID, Username: user.Username, Email: user.Email}, nil
}

func (s *userService) Login(req *dto.LoginUserRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.generateJWT(user)
	if err != nil {
		return nil, err
	}

	userResponse := dto.UserResponse{ID: user.ID, Username: user.Username, Email: user.Email}
	return &dto.LoginResponse{User: userResponse, Token: token}, nil
}

func (s *userService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
