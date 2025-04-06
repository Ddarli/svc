package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type (
	User struct {
		ID        uuid.UUID
		PKey      string
		Username  string
		Password  string
		Email     string
		Phone     string
		Role      int
		CreatedAt time.Time
	}

	AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}
)

func RequestToUser(request RegisterRequest) User {
	return User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Phone:    request.Phone,
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
