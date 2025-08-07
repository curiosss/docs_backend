package middleware

import (
	"docs-notify/internal/models"
	"docs-notify/internal/server"
	"docs-notify/internal/utils/exceptions"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(server *server.Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenHeader := c.Request().Header.Get("Authorization")
			if tokenHeader == "" {
				return exceptions.ErrBadRequest
			}
			claims, err := validateToken(tokenHeader, server.Config.JWTSecret)
			if err != nil {
				return exceptions.ErrUnauthorized
			}
			if claims["user_id"] != nil {
				var userId uint = uint(claims["user_id"].(float64))
				c.Set("user_id", userId)
				user, err := GetUserById(userId, server)
				if err != nil {
					return exceptions.ErrBadRequest
				}
				c.Set("user", user)
			}

			return next(c)
		}
	}
}

func validateToken(tokenHeader string, key string) (jwt.MapClaims, error) {
	parts := strings.Split(tokenHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, exceptions.ErrBadRequest
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exceptions.ErrBadRequest
		}
		// Return secret key used for signing
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, exceptions.ErrUnauthorized
	}

	return claims, nil
}

func GetUserById(userId uint, server *server.Server) (user *models.User, err error) {
	var u models.User
	if err := server.Database.First(&u, userId).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
