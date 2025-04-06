package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func generateToken(_ context.Context, username, secretKey string, expiredHours int) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(expiredHours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
