package controllers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"log"
	"net/http"
	"sss/internal/models"
	"sss/pkg/lib"
	"time"

	"sss/internal/models/orders"
	"strconv"
)

func GetOrders(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	limit, offset := lib.GetLimOff(c)

	q := `SELECT weight, regions, delivery_hours FROM orders LIMIT $1 OFFSET $2`
	rows, err := db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return http.StatusBadRequest, models.BadRequestResponse{}
	}
	defer rows.Close()

	listOrders := make([]orders.OrderDto, 0)
	for rows.Next() {
		tmp := orders.OrderDto{}
		err = rows.Scan(&tmp.Weight, &tmp.Regions, pq.Array(&tmp.DeliveryHours))
		if err != nil {
			return http.StatusBadRequest, models.BadRequestResponse{}
		}
		listOrders = append(listOrders, tmp)
	}

	if len(listOrders) == 0 {
		return http.StatusOK, []orders.OrderDto{}
	}

	return http.StatusOK, listOrders
}

func GetIdOrder(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	tmpId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return http.StatusBadRequest, models.BadRequestResponse{}
	}

	order := orders.OrderDto{}
	order.OrderId = int64(tmpId)

	q := `SELECT weight, regions, delivery_hours FROM orders WHERE order_id = $1`
	err = db.QueryRowContext(ctx, q, tmpId).Scan(&order.Weight, &order.Regions, &order.DeliveryHours)
	if err != nil {
		return http.StatusNotFound, models.BadRequestResponse{}
	}

	return http.StatusOK, order
}

func CreateOrders(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	orderReq := orders.CreateOrderRequest{}
	orderRes := make([]orders.OrderDto, 0)

	err := c.Bind(&orderReq)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, models.InternalServerErrorResponse{}
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO orders (weight, regions, delivery_hours, cost) VALUES ($1, $2, $3, $4) RETURNING order_id`)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, models.InternalServerErrorResponse{}
	}
	defer stmt.Close()

	for _, v := range orderReq.Orders {
		v1 := v
		var orderId int
		err := stmt.QueryRowContext(ctx, v1.Weight, v1.Regions, pq.Array(v1.DeliveryHours), v1.Cost).Scan(&orderId)
		if err != nil {
			tx.Rollback()
			return http.StatusBadRequest, models.BadRequestResponse{}
		}
		var tmp orders.OrderDto
		tmp.OrderId = int64(orderId)
		tmp.Regions = &v1.Regions
		tmp.Weight = &v1.Weight
		tmp.Cost = &v1.Cost
		tmp.DeliveryHours = &v1.DeliveryHours
		orderRes = append(orderRes, tmp)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, models.InternalServerErrorResponse{}
	}

	return http.StatusOK, orderRes
}

func CreateOrdersComplete(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	orderReq := orders.CompleteOrderRequestDto{}
	//orderRes := make([]dto.CompleteOrder, 0)
	err := c.Bind(&orderReq)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, models.InternalServerErrorResponse{}
	}

	stmt, err := tx.PrepareContext(ctx, `SELECT completed_time FROM orders WHERE courier_id = $1 AND order_id = $2`)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, models.InternalServerErrorResponse{}
	}
	defer stmt.Close()

	for _, v := range orderReq.CompleteInfo {
		//v1 := v
		var timeC *time.Time
		err := stmt.QueryRowContext(ctx, v.CourierId, v.OrderId).Scan(&timeC)
		if err != nil {
			tx.Rollback()
			return http.StatusBadRequest, models.BadRequestResponse{}
		}
		//var tmp orders.OrderDto
		//tmp.OrderId = int64(orderId)
		//tmp.Regions = &v1.Regions
		//tmp.Weight = &v1.Weight
		//tmp.Cost = &v1.Cost
		//tmp.DeliveryHours = &v1.DeliveryHours
		//orderRes = append(orderRes, tmp)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, models.InternalServerErrorResponse{}
	}

	return http.StatusOK, nil
}
