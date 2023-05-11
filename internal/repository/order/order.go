package order

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"sss/internal/controllers/dto"
	"sss/internal/entity"
)

type repository struct {
	db *sqlx.DB
}

func NewOrderRero(db *sqlx.DB) Repo {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context, limit, offset int32) ([]entity.Order, error) {
	res := make([]entity.Order, 0)

	q := `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders LIMIT $1 OFFSET $2`
	err := r.db.SelectContext(ctx, &res, q, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return nil, err
	}
	//rows, err := r.db.QueryContext(ctx, q, limit, offset)
	//if err != nil {
	//	return res, err
	//}
	//
	//defer rows.Close()
	//
	//for rows.Next() {
	//	tmp := dto.OrderDto{}
	//	err = rows.Scan(tmp.OrderId, &tmp.Weight, &tmp.Regions, pq.Array(&tmp.DeliveryHours), &tmp.Cost, &tmp.CompletedTime)
	//	if err != nil {
	//		return res, err
	//	}
	//	res = append(res, tmp)
	//}

	return res, nil
}

func (r *repository) GetById(ctx context.Context, id int64) (*entity.Order, error) {
	var order entity.Order
	q := `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
	err := r.db.GetContext(ctx, &order, q, id)
	if err != nil {
		return nil, err
	}
	//err := r.db.QueryRowContext(ctx, q, id).Scan(&order.OrderId, &order.Weight, &order.Regions, pq.Array(&order.DeliveryHours), &order.Cost, &order.CompletedTime)
	//if err != nil {
	//	return nil, err
	//}

	return &order, nil
}

func (r *repository) Create(ctx context.Context, orders *dto.CreateOrderRequest) ([]entity.Order, error) {
	orderRes := make([]entity.Order, 0)

	tx, err := r.db.BeginTxx(ctx, nil)
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
		var tmp entity.Order
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

func (r *repository) Update(ctx context.Context, orders []entity.CompleteOrder) ([]entity.Order, error) {
	res := make([]entity.Order, 0)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var order entity.Order
	var id int64
	for _, v := range orders {
		q := `SELECT cour_id, order_id, weight, regions, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
		err = tx.QueryRowContext(ctx, q, v.OrderId).
			Scan(&id, &order.OrderId, &order.Weight, &order.Regions, &order.DeliveryHours, &order.Cost, &order.CompletedTime)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if id == 0 || id != v.CourierId {
			tx.Rollback()
			return nil, errors.New("id invalidate")
		}

		if order.CompletedTime == nil {
			tmp := v.CompleteTime
			order.CompletedTime = &tmp
			q = `UPDATE orders SET completed_time = $1 WHERE cour_id = $2 AND order_id = $3`
			_, err = tx.ExecContext(ctx, q, v.CompleteTime, v.CourierId, v.OrderId)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		res = append(res, order)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res, nil
}
