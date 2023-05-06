package controllers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"log"
	"net/http"
	"sss/internal/models"
	couriers2 "sss/internal/models/couriers"
	"sss/pkg/lib"
	"strconv"
)

func GetCouriers(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	limit, offset := lib.GetLimOff(c)

	q := `SELECT * FROM couriers LIMIT $1 OFFSET $2`
	rows, err := db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return http.StatusBadRequest, models.BadRequestResponse{}
	}
	defer rows.Close()

	ordersRes := couriers2.GetCouriersResponse{}
	ordersRes.Couriers = make([]couriers2.CourierDto, 0)
	ordersRes.Offset = offset
	ordersRes.Limit = limit
	for rows.Next() {
		tmp := couriers2.CourierDto{}
		err = rows.Scan(&tmp.CourierId, &tmp.CourierType, pq.Array(&tmp.Regions), pq.Array(&tmp.WorkingHours))
		if err != nil {
			return http.StatusBadRequest, models.BadRequestResponse{}
		}
		ordersRes.Couriers = append(ordersRes.Couriers, tmp)
	}

	if len(ordersRes.Couriers) == 0 {
		return http.StatusOK, []couriers2.CourierDto{}
	}
	return http.StatusOK, ordersRes

}

func GetIdCourier(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	tmpId, err := strconv.Atoi(c.Param("courier_id"))
	if err != nil {
		return http.StatusBadRequest, models.BadRequestResponse{}
	}

	courOb := couriers2.CourierDto{}
	courOb.CourierId = int64(tmpId)

	q := `SELECT courier_type, regions, working_hours FROM couriers WHERE courier_id = $1`
	err = db.QueryRowContext(ctx, q, tmpId).Scan(&courOb.CourierType, &courOb.Regions, &courOb.WorkingHours)
	if err != nil {
		return http.StatusNotFound, models.NotFoundResponse{}
	}

	return http.StatusOK, courOb
}

func CreateCouriers(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	courierReq := couriers2.CreateCourierRequest{}
	err := c.Bind(&courierReq)
	if err != nil {
		log.Fatal(err)
	}

	courierRes := couriers2.CreateCouriersResponse{}
	courierRes.Couriers = make([]couriers2.CourierDto, 0, len(courierReq.Couriers))
	for _, v := range courierReq.Couriers {
		q := `INSERT INTO couriers (courier_type, regions, working_hours) VALUES ($1, $2, $3)`
		_, err := db.ExecContext(ctx, q, v.CourierType, pq.Array(v.Regions), pq.Array(v.WorkingHours))
		if err != nil {
			return http.StatusBadRequest, models.BadRequestResponse{}
		}
		var tmp couriers2.CourierDto
		tmp.Regions = v.Regions
		tmp.WorkingHours = v.WorkingHours
		tmp.CourierType = v.CourierType
		courierRes.Couriers = append(courierRes.Couriers, tmp)
	}
	return http.StatusOK, courierRes
}
