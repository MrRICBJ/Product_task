package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"net/http"
	"sss/internal/apperror"
	"sss/internal/courier"
	"sss/internal/order"
	"time"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) courier.Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context, limit, offset int32) (int, interface{}) {
	q := `SELECT * FROM couriers LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return http.StatusBadRequest, apperror.BadRequestResponse{}
	}
	defer rows.Close()

	res := courier.GetCouriersResponse{}
	res.Couriers = make([]courier.Courier, 0)
	res.Offset = offset
	res.Limit = limit
	for rows.Next() {
		tmp := courier.Courier{}
		err = rows.Scan(&tmp.CourierId, &tmp.CourierType, pq.Array(&tmp.Regions), pq.Array(&tmp.WorkingHours))
		if err != nil {
			return http.StatusBadRequest, apperror.BadRequestResponse{}
		}
		res.Couriers = append(res.Couriers, tmp)
	}

	if len(res.Couriers) == 0 {
		return http.StatusOK, []courier.Courier{}
	}
	return http.StatusOK, res
}

func (r *repository) GetMetaInf(ctx context.Context, id int, startDate, endDate time.Time) (int, interface{}) {
	orders := r.getCompletedOrdersForCourier(ctx, id, startDate, endDate)
	//if err != nil {
	//	return 0, err
	//}
	if len(orders) == 0 {
		return 0, nil
	}

	//start := startDate.Format("2006-01-02 00:00:00")
	//end := endDate.Format("2006-01-02 23:59:59")
	//earnings := calculateEarnings(orders, )
	//rating := calculateRating(orders, courierType, startDate, endDate)
	return http.StatusOK, nil
}

//func calculateEarnings(orders []order.Order, courierType string) float64 {
//	var earnings float64
//	for _, order := range orders {
//		earnings += order.Cost * getCoefficient(courierType)
//	}
//	return earnings
//}

func getCoefficient(courierType string) float64 {
	switch courierType {
	case "foot":
		return 2
	case "bike":
		return 3
	case "car":
		return 4
	default:
		return 0
	}
}

func (r *repository) GetById(ctx context.Context, id int) (int, interface{}) {
	courOb := courier.Courier{}
	courOb.CourierId = int64(id)

	q := `SELECT courier_type, regions, working_hours FROM couriers WHERE courier_id = $1`
	err := r.db.QueryRowContext(ctx, q, id).Scan(&courOb.CourierType, &courOb.Regions, &courOb.WorkingHours)
	if err != nil {
		return http.StatusNotFound, apperror.NotFoundResponse{}
	}

	return http.StatusOK, courOb
}

func (r *repository) Create(ctx context.Context, cour *courier.CreateCourierRequest) (int, interface{}) {
	courierRes := courier.CreateCouriersResponse{}
	courierRes.Couriers = make([]courier.Courier, 0, len(cour.Couriers))

	for _, v := range cour.Couriers {
		q := `INSERT INTO couriers (courier_type, regions, working_hours) VALUES ($1, $2, $3)`
		_, err := r.db.ExecContext(ctx, q, v.CourierType, pq.Array(v.Regions), pq.Array(v.WorkingHours))
		if err != nil {
			return http.StatusBadRequest, apperror.BadRequestResponse{}
		}
		var tmp courier.Courier
		tmp.Regions = v.Regions
		tmp.WorkingHours = v.WorkingHours
		tmp.CourierType = v.CourierType
		courierRes.Couriers = append(courierRes.Couriers, tmp)
	}
	return http.StatusOK, courierRes
}

func (r *repository) getCompletedOrdersForCourier(ctx context.Context, id int, startDate, endDate time.Time) []order.Order {
	//q := `SELECT * FROM orders WHERE id = $1 AND complete_time IS NOT NULL`
	q := `SELECT * 
FROM orders 
WHERE cour_id = 1 
AND completed_time IS NOT NULL 
AND completed_time BETWEEN '2023-05-02 00:00:00' AND '2024-01-06 23:59:59'`
	var count int
	err := r.db.QueryRowContext(ctx, q, id).Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		//return http.StatusBadRequest, err ----!!!!!!!!!!!!
	}
	return []order.Order{}
}
