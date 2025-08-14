package jwtutils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId uint, jwtSecret string) (string, error) {

	// Create claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 365 * 100).Unix(), // 100 years
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate access token: " + err.Error())
	}

	return accessToken, nil
}
