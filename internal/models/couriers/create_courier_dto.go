package couriers

type CreateCourierDto struct {
	CourierType string `json:"courier_type"`

	Regions []int32 `json:"regions"`

	WorkingHours []string `json:"working_hours"`
}
