package handlers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"sss/internal/controllers"
)

func GetOrders(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := controllers.GetOrders(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func GetIdOrders(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := controllers.GetIdOrder(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func PostOrders(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := controllers.CreateOrders(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func CreateOrdersComplete(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := controllers.CreateOrdersComplete(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func GetCouriers(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := controllers.GetCouriers(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func GetIdCouriers(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := controllers.GetIdCourier(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func PostCouriers(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := controllers.CreateCouriers(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}
