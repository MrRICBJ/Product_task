package models

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"log"
	"net/http"
	"sss/internal/models/dto"
	"strconv"
)

func GetOrders(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 1
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	q := `SELECT weight, regions, delivery_hours FROM orders LIMIT $1 OFFSET $2`
	rows, err := db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			// обработка ситуации, когда запрос не найден
			return http.StatusOK, []dto.OrderDto{}
		}
		return http.StatusBadRequest, dto.BadRequestResponse{}
	}
	defer rows.Close()

	listOrders := make([]dto.OrderDto, 0)
	for rows.Next() {
		tmp := dto.OrderDto{}
		err = rows.Scan(&tmp.Weight, &tmp.Regions, pq.Array(&tmp.DeliveryHours))
		if err != nil {
			return http.StatusBadRequest, dto.BadRequestResponse{}
		}
		listOrders = append(listOrders, tmp)
	}

	return http.StatusOK, listOrders
}

func PostOrders(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	orderReq := dto.CreateOrderRequest{}
	orderRes := make([]dto.OrderDto, 0)
	err := c.Bind(&orderReq)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range orderReq.Orders {
		v1 := v
		q := `INSERT INTO orders (weight, regions, delivery_hours, cost) VALUES ($1, $2, $3, $4)`
		_, err := db.ExecContext(ctx, q, v1.Weight, v1.Regions, pq.Array(v1.DeliveryHours), v1.Cost)
		if err != nil {
			return http.StatusBadRequest, dto.BadRequestResponse{}
		}
		var tmp dto.OrderDto
		tmp.Regions = &v1.Regions
		tmp.Weight = &v1.Weight
		tmp.Cost = &v1.Cost
		tmp.DeliveryHours = &v1.DeliveryHours
		orderRes = append(orderRes, tmp)
	}
	return http.StatusOK, orderRes
}

func PostOrdersComplete(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	orderReq := dto.CompleteOrderRequestDto{}
	//orderRes := make([]dto.CompleteOrder, 0)
	err := c.Bind(&orderReq)
	if err != nil {
		log.Fatal(err)
	}

}
