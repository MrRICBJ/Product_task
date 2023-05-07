package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"sss/internal/adapters/courier"
	"sss/internal/adapters/order"
	"sss/internal/config"
	dbCour "sss/internal/courier/db"
	dbOrder "sss/internal/order/db"
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

	repositoryOrder := dbOrder.New(db)
	repositoryCour := dbCour.New(db)

	handlerOrder := order.NewHandler(repositoryOrder)
	handlerCour := courier.NewHandler(repositoryCour)

	e := echo.New()

	handlerOrder.Register(e)
	handlerCour.Register(e)

	e.Logger.Fatal(e.Start(":8000"))

}
