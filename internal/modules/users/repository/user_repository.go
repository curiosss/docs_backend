package repository

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/users/dto"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Login(loginDto *dto.UserLoginDto) (*models.User, error) {
	var existing models.User

	// Find user by username
	if err := r.db.Where("username = ?", loginDto.Username).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username")
		}
		return nil, err
	}

	if existing.Password != loginDto.Password {
		return nil, errors.New("invalid password")
	}

	return &existing, nil
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) List(page, limit int) ([]models.User, error) {
	var users []models.User
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetAll() (*models.User, error) {
	var users models.User
	err := r.db.Find(&users).Error
	return &users, err
}

func (r *UserRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
