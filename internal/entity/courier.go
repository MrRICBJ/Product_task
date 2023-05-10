package entity

//import (
//	"database/sql/driver"
//	"fmt"
//	"strings"
//)
//
//type Courier struct {
//	CourierId int64 `json:"courier_id" db:"courier_id"`
//
//	CourierType string `json:"courier_type" db:"courier_type"`
//
//	Regions []int32 `json:"regions" db:"regions"`
//
//	WorkingHours []string `json:"working_hours" db:"working_hours"`
//}
//
//type GetCourierMetaInfoResponse struct {
//	CourierId int64 `json:"courier_id,omitempty" db:"courier_id"`
//
//	CourierType string `json:"courier_type,omitempty" db:"courier_type"`
//
//	Regions []int32 `json:"regions,omitempty" db:"regions,omitempty"`
//
//	WorkingHours workingHours `json:"working_hours,omitempty" db:"working_hours"`
//
//	Rating int32 `json:"rating,omitempty,omitempty"`
//
//	Earnings int32 `json:"earnings,omitempty,omitempty"`
//}
//
//type workingHours []string
//
//func (dh *workingHours) Scan(value interface{}) error {
//	if value == nil {
//		*dh = nil
//		return nil
//	}
//	switch v := value.(type) {
//	case []byte:
//		*dh = workingHours(strings.Split(string(v), ","))
//		return nil
//	case string:
//		*dh = workingHours(strings.Split(v, ","))
//		return nil
//	default:
//		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dh)
//	}
//}
//
//func (dh workingHours) Value() (driver.Value, error) {
//	if dh == nil {
//		return nil, nil
//	}
//	return strings.Join(dh, ","), nil
//}
