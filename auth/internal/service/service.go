package service

import (
	"auth/internal/model"
	"auth/pkg/transport/medicalpb"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type Service struct {
	repo      Repository
	conf      ServiceConf
	bchClient BlockchainProcessor
}

func NewService(repo Repository, conf ServiceConf, bchClient BlockchainProcessor) *Service {
	return &Service{repo: repo, conf: conf, bchClient: bchClient}
}

func (s *Service) Register(ctx context.Context, request model.RegisterRequest) (token string, err error) {
	user := model.RequestToUser(request)
	user.ID = request.ID

	user.Password, err = model.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	res, err := s.bchClient.GenerateNewAccount(ctx, &medicalpb.Empty{})
	if err != nil {
		return "", fmt.Errorf("error btcClient generate pKey: %v", err)
	}
	user.PKey = res.PrivateKey

	err = s.repo.RegisterUser(ctx, user)

	token, err = generateToken(ctx, user.Username, s.conf.SecretKey, s.conf.ExpiredAfterHours)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *Service) Authorize(ctx context.Context, request model.AuthRequest) (string, error) {
	user, err := s.repo.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return "", err
	}

	if model.ComparePasswords(user.Password, request.Password) != nil {
		log.Error("Password is incorrect")
		return "", err
	}

	token, err := generateToken(ctx, user.Username, s.conf.SecretKey, s.conf.ExpiredAfterHours)
	if err != nil {
		log.Error("Error generating token")
		return "", err
	}

	return token, nil
}

func (s *Service) ValidateToken(ctx context.Context, tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Invalid signing method: %v", token.Header["alg"])
		}

		return []byte(s.conf.SecretKey), nil
	})

	if err != nil {
		slog.Error("parsing jwt", "error", err.Error())
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token claims")
	}

	if exp, ok := claims["exp"].(float64); ok {
		expTime := time.Unix(int64(exp), 0)
		if time.Now().After(expTime) {
			return uuid.Nil, fmt.Errorf("token expired at %v", expTime)
		}
	}

	if username, ok := claims["username"].(string); ok {
		if username != "" {
			user, err := s.repo.GetUserByUsername(ctx, username)
			if err != nil {
				return uuid.Nil, fmt.Errorf("invalid username: %v", username)
			}

			return user.ID, nil
		}
	}

	slog.Info("successfully validated token", "username", claims["username"])

	return uuid.Nil, nil
}
