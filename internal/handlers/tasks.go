package handlers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"sss/internal/models"
)

func GetOrders(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := models.GetOrders(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func PostOrders(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := models.PostOrders(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}

func PostOrdersComplete(ctx context.Context, db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode, o := models.PostOrdersComplete(ctx, c, db)
		return c.JSON(statusCode, o)
	}
}
