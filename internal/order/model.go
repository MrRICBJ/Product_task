package order

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Order struct {
	OrderId int64 `json:"order_id,omitempty" db:"order_id"`

	CourierId *int64 `json:"courier_id,omitempty"`

	Weight *float32 `json:"weight,omitempty" db:"weight"`

	Regions *int32 `json:"regions,omitempty" db:"regions"`

	DeliveryHours DeliveryHours `json:"delivery_hours,omitempty" db:"delivery_hours"`

	Cost *int32 `json:"cost,omitempty" db:"cost"`

	CompletedTime *time.Time `json:"completed_time,omitempty" db:"completed_time"`
}

type DeliveryHours []string

func (dh *DeliveryHours) Scan(value interface{}) error {
	if value == nil {
		*dh = nil
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*dh = DeliveryHours(strings.Split(string(v), ","))
		return nil
	case string:
		*dh = DeliveryHours(strings.Split(v, ","))
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dh)
	}
}

func (dh DeliveryHours) Value() (driver.Value, error) {
	if dh == nil {
		return nil, nil
	}
	return strings.Join(dh, ","), nil
}
