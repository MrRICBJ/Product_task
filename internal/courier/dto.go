package courier

type CreateCourierDto struct {
	CourierType string `json:"courier_type"`

	Regions []int32 `json:"regions"`

	WorkingHours []string `json:"working_hours"`
}

type CreateCourierRequest struct {
	Couriers []CreateCourierDto `json:"couriers"`
}

type CreateCouriersResponse struct {
	Couriers []Courier `json:"couriers"`
}

type GetCouriersResponse struct {
	Couriers []Courier `json:"couriers"`

	Limit int32 `json:"limit"`

	Offset int32 `json:"offset"`
}
