package couriers

type CourierDto struct {
	CourierId int64 `json:"courier_id"`

	CourierType string `json:"courier_type"`

	Regions []int32 `json:"regions"`

	WorkingHours []string `json:"working_hours"`
}
