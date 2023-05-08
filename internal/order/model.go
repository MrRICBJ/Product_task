package order

import (
	"time"
)

type Order struct {
	OrderId int64 `json:"order_id,omitempty"`

	CourierId *int64 `json:"courier_id,omitempty"`

	Weight *float32 `json:"weight,omitempty"`

	Regions *int32 `json:"regions,omitempty"`

	DeliveryHours []string `json:"delivery_hours,omitempty"`

	Cost *int32 `json:"cost,omitempty"`

	CompletedTime *time.Time `json:"completed_time,omitempty"`
}
