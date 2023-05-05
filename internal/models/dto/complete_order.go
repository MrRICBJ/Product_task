package dto

import (
	"time"
)

type CompleteOrder struct {
	CourierId int64 `json:"courier_id"`

	OrderId int64 `json:"order_id"`

	CompleteTime time.Time `json:"complete_time"`
}
