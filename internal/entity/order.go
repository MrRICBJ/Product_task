package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Order struct {
	OrderId int64 `json:"order_id" db:"order_id"`

	CourierId *int64 `json:"courier_id,omitempty"`

	Weight float64 `json:"weight" db:"weight"`

	Regions int32 `json:"regions" db:"regions"`

	DeliveryHours deliveryHours `json:"delivery_hours" db:"delivery_hours"`

	Cost int32 `json:"cost" db:"cost"`

	CompletedTime *time.Time `json:"completed_time,omitempty" db:"completed_time"`
}

type deliveryHours []string

func (dh *deliveryHours) Scan(value interface{}) error {
	if value == nil {
		*dh = nil
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*dh = deliveryHours(strings.Split(string(v), ","))
		return nil
	case string:
		*dh = deliveryHours(strings.Split(v, ","))
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dh)
	}
}

func (dh deliveryHours) Value() (driver.Value, error) {
	if dh == nil {
		return nil, nil
	}
	return strings.Join(dh, ","), nil
}
