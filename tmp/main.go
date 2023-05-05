package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"sss/internal"
	"sss/internal/config"
	"sss/pkg/client"
)

func main() {
	//if err := config.InitConfig(); err != nil {
	//	log.Fatalf("error initializing configs: %s\n", err.Error())
	//}

	cfg := config.New()

	ctx, _ := context.WithCancel(context.Background())

	db := client.New(ctx, cfg)
	defer db.Close()

	e := setupServer(ctx, db)
	e.Logger.Fatal(e.Start(":8000"))

}

func setupServer(ctx context.Context, db *sqlx.DB) *echo.Echo {
	e := echo.New()
	internal.SetupRoutes(ctx, e, db)
	return e
}
