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

func (s *UserService) ChangePassword(pwdUpdateDto *dto.UserPwdUpdateDto, userId uint) (*models.User, error) {
	// user, err := s.userRepository.ChangePassword(pwdUpdateDto)
	// if err != nil {
	// 	return nil, err
	// }
	return s.userRepository.ChangePassword(pwdUpdateDto, userId)
}
func (s *UserService) CreateUser(userCreateDto dto.UserCreateDto) (*models.User, error) {
	user := &models.User{
		Username: userCreateDto.Username,
		Password: userCreateDto.Password,
		Role:     userCreateDto.Role,
		Note:     userCreateDto.Note,
	}

	createdUser, err := s.userRepository.Create(user, s.config.JWTSecret)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) UpdateUser(userUpdateDto dto.UserUpdateDto) (*models.User, error) {
	user := &models.User{
		ID:       userUpdateDto.Id,
		Username: userUpdateDto.Username,
		Password: userUpdateDto.Password,
		Role:     userUpdateDto.Role,
		Note:     userUpdateDto.Note,
	}
	createdUser, err := s.userRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) DeleteUser(userId uint) error {
	return s.userRepository.Delete(userId)
}

func (s *UserService) GetUsers() ([]dto.UserResponseDto, error) {
	return s.userRepository.GetAll()
}

func (s *UserService) GetUsersPublic() ([]dto.UserPublicResponseDto, error) {
	return s.userRepository.GetAllPublic()
}
