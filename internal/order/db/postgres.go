package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"net/http"
	"sss/internal/apperror"
	"sss/internal/order"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) order.Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context, limit, offset int32) (int, interface{}) {
	res := make([]order.Order, 0)

	q := `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders LIMIT $1 OFFSET $2`
	err := r.db.SelectContext(ctx, &res, q, limit, offset)
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}

	return http.StatusOK, res
}

func (r *repository) GetById(ctx context.Context, id int) (int, interface{}) {
	order := order.Order{}
	order.OrderId = int64(id)
	q := `SELECT weight, regions, delivery_hours FROM orders WHERE order_id = $1`
	err := r.db.QueryRowContext(ctx, q, id).Scan(&order.Weight, &order.Regions, pq.Array(&order.DeliveryHours))
	if err != nil {
		return http.StatusNotFound, apperror.NotFoundResponse{}
	}

	return http.StatusOK, order
}

func (r *repository) Create(ctx context.Context, orders *order.CreateOrderRequest) (int, interface{}) {
	orderRes := make([]order.Order, 0)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO orders (weight, regions, delivery_hours, cost) VALUES ($1, $2, $3, $4) RETURNING order_id`)
	if err != nil {
		tx.Rollback()
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}
	defer stmt.Close()

	for _, v := range orders.Orders {
		v1 := v
		var orderId int
		err := stmt.QueryRowContext(ctx, v1.Weight, v1.Regions, pq.Array(v1.DeliveryHours), v1.Cost).Scan(&orderId)
		if err != nil {
			tx.Rollback()
			return http.StatusBadRequest, apperror.BadRequestResponse{}
		}
		var tmp order.Order
		tmp.OrderId = int64(orderId)
		tmp.Regions = v1.Regions
		tmp.Weight = v1.Weight
		tmp.Cost = v1.Cost
		tmp.DeliveryHours = v1.DeliveryHours ///////////измен
		orderRes = append(orderRes, tmp)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}

	return http.StatusOK, orderRes
}

func (r *repository) Update(ctx context.Context, orders *order.CompleteOrderRequestDto) (int, interface{}) {
	orderRes := make([]order.Order, 0)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}

	for _, v := range orders.CompleteInfo {
		q := `SELECT courier_id, order_id, completed_time FROM orders WHERE order_id = $1`
		var order order.Order
		err = tx.QueryRowContext(ctx, q, v.OrderId).Scan(&order.CourierId, &order.OrderId, &order.CompletedTime)
		if err != nil {
			tx.Rollback()
			return http.StatusBadRequest, apperror.BadRequestResponse{}
		}

		if order.CourierId == nil || *order.CourierId != v.CourierId {
			tx.Rollback()
			return http.StatusBadRequest, apperror.BadRequestResponse{}
		}

		q = `UPDATE orders SET completed_time = $1 WHERE courier_id = $2 AND order_id = $3`
		//if order.CompletedTime == nil {----
		//	tx.Exec(q, v.CompleteTime, v.CourierId, v.OrderId)----
		//}-----
		order.CourierId = nil
		//order.CompletedTime = nil------
		orderRes = append(orderRes, order)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}

	return http.StatusOK, orderRes
}
