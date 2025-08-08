package database

import (
	"docs-notify/internal/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		// &models.Category{},
		// &models.Doc{},
		// &models.DocUser{},
		// &models.Action{},
		// &models.File{},
	)
	return err
}

func SeedConstants(db *gorm.DB, jwtSecret string) error {
	err := seedAdmin(db, jwtSecret)
	return err
}

func seedAdmin(db *gorm.DB, jwtSecret string) error {

	var count int64
	if err := db.Model(&models.User{}).Where("role = ?", "admin").Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil // Already has an admin, do nothing
	}

	admin := models.User{
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		return errors.New("failed to create default admin: " + err.Error())
	}

	// Create claims
	claims := jwt.MapClaims{
		"user_id":  admin.ID,
		"username": admin.Username,
		"role":     admin.Role,
		"exp":      time.Now().Add(time.Hour * 24 * 365 * 100).Unix(), // 100 years
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return errors.New("failed to generate access token: " + err.Error())
	}

	// Update the admin with the access token
	admin.AccessToken = accessToken
	if err := db.Save(&admin).Error; err != nil {
		return errors.New("failed to update admin with token: " + err.Error())
	}

	return nil
}
