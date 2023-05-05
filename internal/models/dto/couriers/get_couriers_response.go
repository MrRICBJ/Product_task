package couriers

type GetCouriersResponse struct {
	Couriers []CourierDto `json:"couriers"`

	Limit int32 `json:"limit"`

	Offset int32 `json:"offset"`
}
