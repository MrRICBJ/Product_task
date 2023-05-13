package entity

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

//type Order struct {
//	OrderId       int64      `json:"order_id" db:"order_id"`
//	Weight        float32    `json:"weight" db:"weight"`
//	Regions       int32      `json:"regions" db:"regions"`
//	DeliveryHours []string   `json:"delivery_hours" db:"delivery_hours"`
//	Cost          int32      `json:"cost" db:"cost"`
//	CompletedTime *time.Time `json:"completed_time,omitempty" db:"completed_time"`
//}
//
//type CompleteOrder struct {
//	CourierId int64 `json:"courier_id"`
//
//	OrderId int64 `json:"order_id"`
//
//	CompleteTime time.Time `json:"complete_time"`
//}

type Order struct {
	OrderId       int64      `json:"order_id" db:"order_id"`
	Weight        float32    `json:"weight" db:"weight"`
	Regions       int32      `json:"regions" db:"regions"`
	DeliveryHours sliceStr   `json:"delivery_hours" db:"delivery_hours"`
	Cost          int32      `json:"cost" db:"cost"`
	CompletedTime *time.Time `json:"completed_time,omitempty" db:"completed_time"`
}

type sliceStr []string

func (dh *sliceStr) Scan(value interface{}) error {
	if value == nil {
		*dh = nil
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*dh = sliceStr(strings.Split(string(v), ","))
		return nil
	case string:
		*dh = sliceStr(strings.Split(v, ","))
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dh)
	}
}

func (dh sliceStr) Value() (driver.Value, error) {
	if dh == nil {
		return nil, nil
	}
	return strings.Join(dh, ","), nil
}
