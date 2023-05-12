package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sss/config"
)

func NewPostgresDB(cfg config.PostgresConfig) (*sqlx.DB, error) {
	q := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	//db, err := sqlx.Open("postgres", q)

	db, err := sqlx.Connect("postgres", q)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return db, nil

	//fmt.Println(cfg, err)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = db.Ping()
	//if err != nil {
	//	return nil, err
	//}
	//
	//return db, nil
}
