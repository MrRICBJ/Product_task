package orders

import (
	"time"
)

type OrderDto struct {
	OrderId int64 `json:"order_id,omitempty"`

	Weight *float32 `json:"weight,omitempty"`

	Regions *int32 `json:"regions,omitempty"`

	DeliveryHours *[]string `json:"delivery_hours,omitempty"`

	Cost *int32 `json:"cost,omitempty"`

	CompletedTime *time.Time `json:"completed_time,omitempty"`
}
