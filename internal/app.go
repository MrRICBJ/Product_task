package internal

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"sss/internal/handlers"
)

func SetupRoutes(ctx context.Context, e *echo.Echo, db *sqlx.DB) {
	e.GET("/orders", handlers.GetOrders(ctx, db))
	e.POST("/orders", handlers.PostOrders(ctx, db))
	e.POST("/orders/complete", handlers.PostOrdersComplete(ctx, db))
}