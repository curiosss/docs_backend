package database

import (
	"context"
	"docs-notify/internal/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func Connect(dsn string, cfg *config.Config) (*gorm.DB, error) {
	database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Warn),
			NowFunc: func() time.Time {
				utc, _ := time.LoadLocation("")
				return time.Now().In(utc)
			},
			PrepareStmt:     false,
			CreateBatchSize: 100,
			Dialector:       postgres.New(postgres.Config{DSN: dsn}),
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(40)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Minute * 15)

	_, err = sqlDB.Conn(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error connection database: %v", err)
	}

	if !cfg.DisableAutoMigration {
		// Миграция схемы
		err := RunMigrations(database)
		if err != nil {
			return nil, fmt.Errorf("error migration database: %v", err)
		}
	}

	err = SeedConstants(database, cfg.JWTSecret)
	if err != nil {
		return nil, fmt.Errorf("error seeding constants: %v", err)
	}

	log.Println("Connected to PostgresSql")
	return database, nil
}
