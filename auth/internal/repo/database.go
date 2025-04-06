package repo

import (
	"auth/conf"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func ConnectDB(ctx context.Context, cfg conf.Configuration) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host,
		cfg.Database.Port, cfg.Database.Name)

	var err error

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Connected to database")

	return db
}
