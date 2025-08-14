package database

import (
	"docs-notify/internal/models"
	jwtutils "docs-notify/internal/utils/jwt_utils"
	"errors"

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

	accessToken, err := jwtutils.GenerateToken(admin.ID, jwtSecret)
	if err != nil {
		return err
	}

	// Update the admin with the access token
	admin.AccessToken = accessToken
	if err := db.Save(&admin).Error; err != nil {
		return errors.New("failed to update admin with token: " + err.Error())
	}

	return nil
}
