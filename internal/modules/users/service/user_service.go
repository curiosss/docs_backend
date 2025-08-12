package service

import (
	"docs-notify/internal/config"
	"docs-notify/internal/models"
	"docs-notify/internal/modules/users/dto"
	"docs-notify/internal/modules/users/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
	config         *config.Config
}

func NewUserService(userRepository *repository.UserRepository, cfg *config.Config) *UserService {
	return &UserService{userRepository: userRepository, config: cfg}
}

func (s *UserService) Login(loginDto dto.UserLoginDto) (*models.User, error) {
	user, err := s.userRepository.Login(&loginDto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) ChangeUsername(loginDto *dto.UserLoginDto, userId uint) (*models.User, error) {
	user, err := s.userRepository.ChangeUsername(loginDto, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) ChangePassword(loginDto *dto.UserLoginDto) (*models.User, error) {
	user, err := s.userRepository.ChangePassword(loginDto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 	res, _ := s.userRepository.GetByEmail(user.Email)
// 	if res != nil && res.Email == user.Email {
// 		return nil, errors.New("user with this email already exists")
// 	}

// 	return s.userRepository.Create(user)
// }

// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
// 	return nil, err
// }

// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 	"user_id": user.Id,
// 	"exp":     time.Now().Add(time.Minute * 20).Unix(),
// })
// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 	"user_id": user.Id,
// 	"exp":     time.Now().Add(time.Hour * 24 * 10).Unix(),
// })
// tokenString, err := token.SignedString([]byte(s.config.SecretKeyForUser))
// if err != nil {
// 	return nil, err
// }
// refreshString, err := refreshToken.SignedString([]byte(s.config.SecretKeyForRefresh))
// if err != nil {
// 	return nil, err
// }

// return &dto.UserResponseDto{
// 	ID:           user.Id,
// 	Username:     user.Username,
// 	Name:         user.Name,
// 	Email:        user.Email,
// 	AccessToken:  tokenString,
// 	RefreshToken: refreshString,
// }, nil

// }

func (s *UserService) GetAll() (*models.User, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// func (s *UserService) Update(userId uint, updateDto dto.UserUpdateDto) (*models.User, error) {
// 	user, err := s.userRepository.GetById(userId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if updateDto.Username != "" {
// 		user.Username = updateDto.Username
// 	}
// 	if updateDto.Name != "" {
// 		user.Name = updateDto.Name
// 	}
// 	if updateDto.Email != "" {
// 		user.Email = updateDto.Email
// 	}
// 	if updateDto.Password != "" {
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateDto.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			return nil, err
// 		}
// 		user.Password = string(hashedPassword)
// 	}

// 	return s.userRepository.Update(user)
// }
