package repository

import (
	"docs-notify/internal/models"
	"docs-notify/internal/modules/users/dto"
	jwtutils "docs-notify/internal/utils/jwt_utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Login(loginDto *dto.UserLoginDto) (*models.User, error) {
	var existing models.User

	// Find user by username
	if err := r.db.Where("username = ?", loginDto.Username).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Ulanyjy tapylmady")
		}
		return nil, err
	}

	if existing.Password != loginDto.Password {
		return nil, errors.New("Parol dogry däl")
	}
	existing.FcmToken = loginDto.FcmToken
	if err := r.db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *UserRepository) ChangeUsername(loginDto *dto.UserLoginDto, userId uint) (*models.User, error) {
	var existing models.User

	if err := r.db.Where("id = ?", userId).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Ulanyjy tapylmady")
		}
		return nil, err
	}

	if existing.Password != loginDto.Password {
		return nil, errors.New("Parol dogry däl")
	}

	existing.Username = loginDto.Username
	if err := r.db.Save(existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *UserRepository) ChangePassword(pwdUpdateDto *dto.UserPwdUpdateDto, userId uint) (*models.User, error) {
	var existing models.User

	if err := r.db.Where("id = ?", userId).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Ulanyjy tapylmady")
		}
		return nil, err
	}

	if existing.Password != pwdUpdateDto.Password {
		return nil, errors.New("Parol dogry däl")
	}

	existing.Password = pwdUpdateDto.NewPassword
	if err := r.db.Save(existing).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (r *UserRepository) Create(user *models.User, jwtSecret string) (*models.User, error) {
	// Check if username already exists
	var count int64
	if err := r.db.Model(&models.User{}).
		Where("username = ?", user.Username).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("Bu ulanyjy ady ulgama hasaba alnan")
	}

	// Create the user
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	token, err := jwtutils.GenerateToken(user.ID, jwtSecret)
	if err != nil {
		return nil, err
	}
	user.AccessToken = token

	// Save the user with the access token
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	// Check if username is taken by another user
	var existing models.User
	err := r.db.
		Where("username = ? AND id <> ?", user.Username, user.ID).
		First(&existing).Error

	if err == nil {
		return nil, errors.New("ulanyjy ady başgalary tarapyndan ulanylýar")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = r.db.Model(&user).
		Select("username", "password", "role", "note").
		Updates(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (r *UserRepository) GetByID(id uint) (*models.User, error) {
// 	var user models.User
// 	err := r.db.First(&user, id).Error
// 	return &user, err
// }

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	fmt.Println(result.Error)
	if result.RowsAffected > 0 {
		return nil
	} else {
		if result.Error == nil {
			return errors.New("Ulanyjy tapylmady")
		}
		return result.Error
	}
}

// func (r *UserRepository) GetById(id uint) (*models.User, error) {
// 	var user models.User
// 	if err := r.db.First(&user, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

func (r *UserRepository) GetAll() ([]dto.UserResponseDto, error) {
	var responses []dto.UserResponseDto

	// Fetch only necessary fields and order by role
	if err := r.db.Model(&models.User{}).
		Order("role ASC").
		Scan(&responses).Error; err != nil {
		return nil, err
	}
	return responses, nil
}

func (r *UserRepository) GetAllPublic() ([]dto.UserPublicResponseDto, error) {
	var responses []dto.UserPublicResponseDto

	// Fetch only necessary fields and order by role
	if err := r.db.Model(&models.User{}).
		Order("role ASC").
		Scan(&responses).Error; err != nil {
		return nil, err
	}
	return responses, nil
}

func (r *UserRepository) GetAllNotif() ([]dto.UserNotifDto, error) {
	var responses []dto.UserNotifDto

	// Fetch only necessary fields and order by role
	if err := r.db.Model(&models.User{}).
		Order("role ASC").
		Scan(&responses).Error; err != nil {
		return nil, err
	}
	return responses, nil
}

func (r *UserRepository) RemoveFcmToken(userId uint) error {
	if err := r.db.Table("users").
		Where("id = ?", userId).
		Update("fcm_token", "").Error; err != nil {
		return err
	}
	return nil
}
