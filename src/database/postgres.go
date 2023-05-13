package database

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sss/config"
)

func NewPostgresDB(cfg config.PostgresConfig) (*sqlx.DB, error) {
	q := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	//q := "host=db port=5432 user=postgres dbname=postgres password=password sslmode=disable"
	//fmt.Println("+++++++" + q + "++++++++++")
	fmt.Println(q)
	db, err := sqlx.Connect("postgres", q)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	// Создаем экземпляр Migrate
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations/.",
		"postgres", driver)
	if err != nil {
		return nil, err
	}

	// Выполняем миграции вверх
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return db, nil
}
