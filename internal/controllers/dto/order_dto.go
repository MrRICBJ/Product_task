package dto

import (
	"sss/internal/entity"
	"time"
)

type OrderDto struct {
	OrderId       int64      `json:"order_id"`
	Weight        float32    `json:"weight"`
	Regions       int32      `json:"regions"`
	DeliveryHours []string   `json:"delivery_hours"`
	Cost          int32      `json:"cost"`
	CompletedTime *time.Time `json:"completed_time,omitempty"`
}

type CreateOrderDto struct {
	Weight        float32  `json:"weight"`
	Regions       int32    `json:"regions"`
	DeliveryHours []string `json:"delivery_hours"`
	Cost          int32    `json:"cost"`
}

type CreateOrderRequest struct {
	Orders []CreateOrderDto `json:"orders"`
}

type CompleteOrderRequestDto struct {
	CompleteInfo []entity.CompleteOrder `json:"complete_info"`
}
