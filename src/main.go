package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"sss/config"
	v1 "sss/controllers/v1"
	"sss/database"
	"sss/repository/courier"
	"sss/repository/order"
	"sss/server"
	"sss/service"
	"syscall"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s\n", err.Error())
	}

	cfg := config.New()

	db, err := database.NewPostgresDB(cfg.Post)
	if err != nil {
		logrus.Fatalf("db opening error: %s\n", err.Error())
	}

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), app(db)); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func app(db *sqlx.DB) *echo.Echo {

	repositoryOrder := order.NewOrderRero(db)
	repositoryCour := courier.NewCourRepo(db)

	orderService := service.NewOrderService(repositoryOrder)
	courService := service.NewCourService(repositoryCour)

	handlerOrder := v1.NewOrderHandler(orderService)
	handlerCour := v1.NewCourHandler(courService)

	//rate := limiter.Rate{
	//	Limit:  10,
	//	Period: time.Second,
	//}

	e := echo.New()
	//e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	handlerOrder.Register(e)
	handlerCour.Register(e)

	return e
}
