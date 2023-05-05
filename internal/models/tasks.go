package models

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"log"
	"net/http"
	"sss/internal/models/dto"
	"sss/internal/models/dto/couriers"
	"sss/internal/models/dto/orders"
	"strconv"
)

func GetOrders(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	limit, offset := getLimOff(c)

	q := `SELECT weight, regions, delivery_hours FROM orders LIMIT $1 OFFSET $2`
	rows, err := db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return http.StatusBadRequest, dto.BadRequestResponse{}
	}
	defer rows.Close()

	listOrders := make([]orders.OrderDto, 0)
	for rows.Next() {
		tmp := orders.OrderDto{}
		err = rows.Scan(&tmp.Weight, &tmp.Regions, pq.Array(&tmp.DeliveryHours))
		if err != nil {
			return http.StatusBadRequest, dto.BadRequestResponse{}
		}
		listOrders = append(listOrders, tmp)
	}

	if len(listOrders) == 0 {
		return http.StatusOK, []orders.OrderDto{}
	}

	return http.StatusOK, listOrders
}

func GetIdOrders(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	tmpId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return http.StatusBadRequest, dto.BadRequestResponse{}
	}

	order := orders.OrderDto{}
	order.OrderId = int64(tmpId)

	q := `SELECT weight, regions, delivery_hours FROM orders WHERE order_id = $1`
	err = db.QueryRowContext(ctx, q, tmpId).Scan(&order.Weight, &order.Regions, &order.DeliveryHours)
	if err != nil {
		return http.StatusNotFound, dto.BadRequestResponse{}
	}

	return http.StatusOK, order
}

func PostOrders(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	orderReq := orders.CreateOrderRequest{}
	orderRes := make([]orders.OrderDto, 0)
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
		var tmp orders.OrderDto
		tmp.Regions = &v1.Regions
		tmp.Weight = &v1.Weight
		tmp.Cost = &v1.Cost
		tmp.DeliveryHours = &v1.DeliveryHours
		orderRes = append(orderRes, tmp)
	}
	return http.StatusOK, orderRes
}

//func PostOrdersComplete(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
//	orderReq := dto.CompleteOrderRequestDto{}
//	//orderRes := make([]dto.CompleteOrder, 0)
//	err := c.Bind(&orderReq)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}

func GetCouriers(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	limit, offset := getLimOff(c)

	q := `SELECT * FROM couriers LIMIT $1 OFFSET $2`
	rows, err := db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return http.StatusBadRequest, dto.BadRequestResponse{}
	}
	defer rows.Close()

	ordersRes := couriers.GetCouriersResponse{}
	ordersRes.Couriers = make([]couriers.CourierDto, 0)
	ordersRes.Offset = offset
	ordersRes.Limit = limit
	for rows.Next() {
		tmp := couriers.CourierDto{}
		err = rows.Scan(&tmp.CourierId, &tmp.CourierType, pq.Array(&tmp.Regions), pq.Array(&tmp.WorkingHours))
		if err != nil {
			return http.StatusBadRequest, dto.BadRequestResponse{}
		}
		ordersRes.Couriers = append(ordersRes.Couriers, tmp)
	}

	if len(ordersRes.Couriers) == 0 {
		return http.StatusOK, []couriers.CourierDto{}
	}
	return http.StatusOK, ordersRes

}

func GetIdCouriers(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	tmpId, err := strconv.Atoi(c.Param("courier_id"))
	if err != nil {
		return http.StatusBadRequest, dto.BadRequestResponse{}
	}

	courOb := couriers.CourierDto{}
	courOb.CourierId = int64(tmpId)

	q := `SELECT courier_type, regions, working_hours FROM couriers WHERE courier_id = $1`
	err = db.QueryRowContext(ctx, q, tmpId).Scan(&courOb.CourierType, &courOb.Regions, &courOb.WorkingHours)
	if err != nil {
		return http.StatusNotFound, dto.NotFoundResponse{}
	}

	return http.StatusOK, courOb
}

func PostCouriers(ctx context.Context, c echo.Context, db *sqlx.DB) (int, interface{}) {
	courierReq := couriers.CreateCourierRequest{}
	err := c.Bind(&courierReq)
	if err != nil {
		log.Fatal(err)
	}

	courierRes := couriers.CreateCouriersResponse{}
	courierRes.Couriers = make([]couriers.CourierDto, 0, len(courierReq.Couriers))
	for _, v := range courierReq.Couriers {
		q := `INSERT INTO couriers (courier_type, regions, working_hours) VALUES ($1, $2, $3)`
		_, err := db.ExecContext(ctx, q, v.CourierType, pq.Array(v.Regions), pq.Array(v.WorkingHours))
		if err != nil {
			return http.StatusBadRequest, dto.BadRequestResponse{}
		}
		var tmp couriers.CourierDto
		tmp.Regions = v.Regions
		tmp.WorkingHours = v.WorkingHours
		tmp.CourierType = v.CourierType
		courierRes.Couriers = append(courierRes.Couriers, tmp)
	}
	return http.StatusOK, courierRes
}

func getLimOff(c echo.Context) (int32, int32) {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 1
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	return int32(limit), int32(offset)
}
