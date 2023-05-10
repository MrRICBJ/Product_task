package dto

type CourierDto struct {
	CourierId    int64    `json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []int32  `json:"regions"`
	WorkingHours []string `json:"working_hours"`
}

type GetCouriersResponse struct {
	Couriers []CourierDto `json:"couriers"`
	Limit    int32        `json:"limit"`
	Offset   int32        `json:"offset"`
}

type CreateCourierRequest struct {
	Couriers []CreateCourierDto `json:"couriers"`
}

type CreateCourierDto struct {
	CourierType  string   `json:"courier_type"`
	Regions      []int32  `json:"regions"`
	WorkingHours []string `json:"working_hours"`
}

type CreateCouriersResponse struct {
	Couriers []CourierDto `json:"couriers"`
}

type GetCourierMetaInfoResponse struct {
	CourierId    int64    `json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []int32  `json:"regions"`
	WorkingHours []string `json:"working_hours"`
	Rating       int32    `json:"rating,omitempty"`
	Earnings     int32    `json:"earnings,omitempty"`
}
