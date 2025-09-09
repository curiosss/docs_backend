package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort              string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	JWTSecret            string
	DisableAutoMigration bool
}

func LoadConfig() *Config {
	// return &Config{
	// 	AppPort:              "8000",
	// 	DBHost:               "localhost",
	// 	DBPort:               "5432",
	// 	DBUser:               "user",
	// 	DBPassword:           "pwd4docs3",
	// 	DBName:               "docs_notify_db",
	// 	JWTSecret:            "Docs4notifier7",
	// 	DisableAutoMigration: false,
	// }

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load it")
	}

	disableAutoMigration, _ := strconv.ParseBool(getEnv("DISABLE_AUTO_MIGRATION", "false"))
	return &Config{
		AppPort:              getEnv("APP_PORT", "8080"),
		DBHost:               getEnv("DB_HOST", "localhost"),
		DBPort:               getEnv("DB_PORT", "5432"),
		DBUser:               getEnv("DB_USER", "user"),
		DBPassword:           getEnv("DB_PASSWORD", "password"),
		DBName:               getEnv("DB_NAME", "docs_notify_db"),
		JWTSecret:            getEnv("JWT_SECRET", "default_secret"),
		DisableAutoMigration: disableAutoMigration,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("Using fallback for env var %s", key)
	return fallback
}
