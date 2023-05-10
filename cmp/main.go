package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"sss/internal/config"
	"sss/internal/controllers/v1"
	"sss/internal/repository/courier"
	"sss/internal/repository/order"
	"sss/internal/service"
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

	repositoryOrder := order.NewOrderRero(db)
	repositoryCour := courier.NewCourRepo(db)

	orderService := service.NewOrderService(repositoryOrder)
	courService := service.NewCourService(repositoryCour)

	handlerOrder := v1.NewOrderHandler(orderService)
	handlerCour := v1.NewCourHandler(courService)

	e := echo.New()

	handlerOrder.Register(e)
	handlerCour.Register(e)

	e.Logger.Fatal(e.Start(":8000"))

}
