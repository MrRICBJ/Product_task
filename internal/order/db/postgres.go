package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"net/http"
	"sss/internal/apperror"
	"sss/internal/order"
	"time"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) order.Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context, limit, offset int32) (int, interface{}) {
	q := `SELECT weight, regions, delivery_hours FROM orders LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}
	defer rows.Close()

	listOrders := make([]order.Order, 0)
	for rows.Next() {
		tmp := order.Order{}
		err = rows.Scan(&tmp.Weight, &tmp.Regions, pq.Array(&tmp.DeliveryHours))
		if err != nil {
			return http.StatusBadRequest, apperror.BadRequestResponse{}
		}
		listOrders = append(listOrders, tmp)
	}

	if len(listOrders) == 0 {
		return http.StatusOK, []order.Order{}
	}

	return http.StatusOK, listOrders
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
		tmp.Regions = &v1.Regions
		tmp.Weight = &v1.Weight
		tmp.Cost = &v1.Cost
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
	orderRes := make([]int64, 0)
	//orderRes := make([]order.Order, 0)
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}

	stmt, err := tx.PrepareContext(ctx, `SELECT completed_time FROM orders WHERE courier_id = $1 AND order_id = $2`)
	if err != nil {
		tx.Rollback()
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}
	defer stmt.Close()

	for _, v := range orders.CompleteInfo {
		var timeC *time.Time
		err := stmt.QueryRowContext(ctx, v.CourierId, v.OrderId).Scan(&timeC)
		if err != nil {
			tx.Rollback()
			return http.StatusBadRequest, apperror.BadRequestResponse{}
		}
		if timeC == nil {
			r.db.Exec("UPDATE orders SET completed_time = $1 WHERE courier_id = $2 AND order_id = $3", v.CompleteTime, v.CourierId, v.OrderId)
		}
		//var tmp order.Order{}
		//tmp.OrderId = int64(orderId)
		//tmp.Regions = v.
		//tmp.Weight = v.Weight
		//tmp.Cost = v.Cost
		//tmp.DeliveryHours = &v1.DeliveryHours
		//orderRes = append(orderRes, tmp)
		orderRes = append(orderRes, v.OrderId)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}

	return http.StatusOK, orderRes
}
