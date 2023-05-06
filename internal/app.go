package internal

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"sss/internal/handlers"
)

func SetupRoutes(ctx context.Context, e *echo.Echo, db *sqlx.DB) {
	e.GET("/orders", handlers.GetOrders(ctx, db))
	e.GET("/orders/:order_id", handlers.GetIdOrders(ctx, db))
	e.POST("/orders", handlers.PostOrders(ctx, db))
	e.POST("/orders/complete", handlers.CreateOrdersComplete(ctx, db))

	e.GET("/couriers", handlers.GetCouriers(ctx, db))
	e.GET("/couriers/:courier_id", handlers.GetIdCouriers(ctx, db))
	e.POST("/couriers", handlers.PostCouriers(ctx, db))
}
