package repo

import (
	"auth/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}

func (r *repo) RegisterUser(ctx context.Context, user model.User) error {
	_, err := r.db.Exec(ctx, insertNewUser, user.ID, user.PKey, user.Username, user.Password, user.Email, user.Phone, 2)
	if err != nil {
		log.Printf("Error inserting new user: %v", err)
	}

	return err
}

func (r *repo) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User

	err := r.db.QueryRow(ctx, selectUserByUsername, username).Scan(&user.ID, &user.PKey, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Role)
	if err != nil {
		log.Printf("Error getting user by username: %v", err)
	}

	return user, err
}
