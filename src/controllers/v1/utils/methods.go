package utils

import (
	"errors"
	"regexp"
	"sss/controllers/dto"
	"sss/entity"
	"strconv"
)

const (
	foot = "FOOT"
	bike = "BIKE"
	car  = "CAR"
)

func GetLimOff(limStr, offStr string) (int32, int32, error) {
	var limit, offset int
	var err error

	if limStr == "" {
		limit = 1
	} else {
		limit, err = strconv.Atoi(limStr)
		if err != nil {
			return 0, 0, err
		}
	}

	if offStr != "" {
		offset, err = strconv.Atoi(offStr)
		if err != nil {
			return 0, 0, err
		}
	}

	return int32(limit), int32(offset), nil
}

func ConvertOrdersToOrderDtos(orders []entity.Order) []dto.OrderDto {
	orderDtos := make([]dto.OrderDto, len(orders))
	for i, order := range orders {
		orderDtos[i] = dto.OrderDto{
			OrderId:       order.OrderId,
			Weight:        order.Weight,
			Regions:       order.Regions,
			DeliveryHours: []string(order.DeliveryHours),
			Cost:          order.Cost,
			CompletedTime: order.CompletedTime,
		}
	}
	return orderDtos
}

func ConvertToCourierDto(couriers []entity.Courier) []dto.CourierDto {
	courierDtos := make([]dto.CourierDto, 0)
	for _, courier := range couriers {
		courierDto := dto.CourierDto{
			CourierId:    courier.CourierId,
			CourierType:  courier.CourierType,
			Regions:      courier.Regions,
			WorkingHours: courier.WorkingHours,
		}
		courierDtos = append(courierDtos, courierDto)
	}
	return courierDtos
}

func ValidationOrder(o *dto.CreateOrderRequest) error {
	for _, order := range o.Orders {
		if len(order.DeliveryHours) == 0 || order.Cost < 0 || order.Regions < 0 || order.Weight < 0 {
			return errors.New("")
		}
		for _, interval := range order.DeliveryHours {
			valid, err := regexp.MatchString(`^\d{2}:\d{2}-\d{2}:\d{2}$`, interval)
			if err != nil {
				return errors.New("")
			}
			if !valid {
				return errors.New("")
			}
		}
	}
	return nil
}

func ValidationCour(o *dto.CreateCourierRequest) error {
	for _, cour := range o.Couriers {
		if len(cour.WorkingHours) == 0 || len(cour.Regions) == 0 {
			return errors.New("")
		}
		for _, interval := range cour.WorkingHours {
			valid, err := regexp.MatchString(`^\d{2}:\d{2}-\d{2}:\d{2}$`, interval)
			if err != nil || !valid {
				return errors.New("")
			}
		}
		for _, reg := range cour.Regions {
			if reg < 0 {
				return errors.New("")
			}
		}
		if cour.CourierType != foot && cour.CourierType != bike && cour.CourierType != car {
			return errors.New("")
		}
	}
	return nil
}

// { post
//     "orders": [
//     {
//       "weight": 1,
//       "regions": 1,
//       "delivery_hours": [
//         "12:12-12:"
//       ],
//       "cost": 1
//     }
//   ]
// }
