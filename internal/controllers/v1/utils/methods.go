package utils

import (
	"errors"
	"regexp"
	"sss/internal/controllers/dto"
	"sss/internal/entity"
	"strconv"
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
	var courierDtos []dto.CourierDto
	for _, courier := range couriers {
		courierDto := dto.CourierDto{
			CourierId:    courier.CourierId,
			CourierType:  courier.CourierType,
			Regions:      courier.Regions,
			WorkingHours: []string(courier.WorkingHours),
		}
		courierDtos = append(courierDtos, courierDto)
	}
	return courierDtos
}

func Validation(o *dto.CreateOrderRequest) error {
	for _, order := range o.Orders {
		if len(order.DeliveryHours) == 0 {
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
