package order

import "time"

type CompleteOrder struct {
	CourierId int64 `json:"courier_id"`

	OrderId int64 `json:"order_id"`

	CompleteTime time.Time `json:"complete_time"`
}

type CompleteOrderRequestDto struct {
	CompleteInfo []CompleteOrder `json:"complete_info"`
}

type CreateOrderDto struct {
	Weight float64 `json:"weight"`

	Regions int32 `json:"regions"`

	DeliveryHours []string `json:"delivery_hours"`

	Cost int32 `json:"cost"`
}

type CreateOrderRequest struct {
	Orders []CreateOrderDto `json:"orders"`
}
