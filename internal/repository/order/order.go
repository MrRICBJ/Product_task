package order

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"sss/internal/controllers/dto"
)

type repository struct {
	db *sqlx.DB
}

func NewOrderRero(db *sqlx.DB) Repo {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context, limit, offset int32) ([]dto.OrderDto, error) {
	res := make([]dto.OrderDto, 0)

	q := `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		tmp := dto.OrderDto{}
		err = rows.Scan(tmp.OrderId, &tmp.Weight, &tmp.Regions, pq.Array(&tmp.DeliveryHours), &tmp.Cost, &tmp.CompletedTime)
		if err != nil {
			return res, err
		}
		res = append(res, tmp)
	}

	return res, err
}

func (r *repository) GetById(ctx context.Context, id int64) (dto.OrderDto, error) {
	order := dto.OrderDto{}
	order.OrderId = id
	q := `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
	err := r.db.QueryRowContext(ctx, q, id).Scan(&order.OrderId, &order.Weight, &order.Regions, pq.Array(&order.DeliveryHours), &order.Cost, &order.CompletedTime)
	if err != nil {
		return dto.OrderDto{}, err
	}

	return order, nil
}

func (r *repository) Create(ctx context.Context, orders dto.CreateOrderRequest) ([]dto.OrderDto, error) {
	orderRes := make([]dto.OrderDto, 0)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	q := `INSERT INTO orders (weight, regions, delivery_hours, cost) VALUES ($1, $2, $3, $4) RETURNING order_id`
	for _, v := range orders.Orders {
		var orderId int64
		err := tx.QueryRowContext(ctx, q, v.Weight, v.Regions, pq.Array(v.DeliveryHours), v.Cost).Scan(&orderId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		var tmp dto.OrderDto
		tmp.OrderId = orderId
		tmp.Regions = v.Regions
		tmp.Weight = v.Weight
		tmp.Cost = v.Cost
		tmp.DeliveryHours = v.DeliveryHours
		orderRes = append(orderRes, tmp)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return orderRes, err
}

func (r *repository) Update(ctx context.Context, orders dto.CompleteOrderRequestDto) ([]dto.OrderDto, error) {
	res := make([]dto.OrderDto, 0)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	q := `SELECT courier_id, order_id, weight, regions, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
	var order dto.OrderDto
	var id int64
	for _, v := range orders.CompleteInfo {
		err = tx.QueryRowContext(ctx, q, v.OrderId).
			Scan(&id, &order.OrderId, &order.Weight, &order.Regions, pq.Array(order.DeliveryHours), &order.Cost, &order.CompletedTime)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if id == 0 || id != v.CourierId {
			tx.Rollback()
			return nil, err
		}

		q = `UPDATE orders SET completed_time = $1 WHERE courier_id = $2 AND order_id = $3`
		_, err = tx.ExecContext(ctx, q, v.OrderId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		res = append(res, order)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res, nil
}
