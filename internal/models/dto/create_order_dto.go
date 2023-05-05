package dto

type CreateOrderDto struct {
	Weight float32 `json:"weight"`

	Regions int32 `json:"regions"`

	DeliveryHours []string `json:"delivery_hours"`

	Cost int32 `json:"cost"`
}
