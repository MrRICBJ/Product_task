package entity

type Courier struct {
	CourierId    int64    `json:"courier_id" db:"courier_id"`
	CourierType  string   `json:"courier_type" db:"courier_type"`
	Regions      []int32  `json:"regions" db:"regions"`
	WorkingHours []string `json:"working_hours" db:"working_hours"`
}
