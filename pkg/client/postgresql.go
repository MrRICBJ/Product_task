package client

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sss/internal/config"
)

func New(ctx context.Context, config config.Config) *sqlx.DB {
	db, err := dbConnect(ctx, config.Post)
	if err != nil {
		log.Fatal("Failed to connect to postgre", err)
	}
	return db

}

func dbConnect(ctx context.Context, cfg config.PostgresConfig) (*sqlx.DB, error) {
	q := "host=localhost port=5432 user=postgres dbname=yandexDB password=2002 sslmode=disable"
	db, err := sqlx.ConnectContext(ctx, "postgres", q)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
