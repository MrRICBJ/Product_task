package utils

import (
	"errors"
	"regexp"
	"sort"
	"sss/controllers/dto"
	"sss/entity"
	"strconv"
)

const (
	foot = "FOOT"
	bike = "BIKE"
	car  = "CAR"
)

type timeInterval struct {
	start string
	end   string
}

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
			DeliveryHours: order.DeliveryHours,
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
		if dublTime(order.DeliveryHours) {
			return errors.New("")
		}
	}
	return nil
}

func ValidationCour(o *dto.CreateCourierRequest) error {
	errChan1 := make(chan error)
	errChan2 := make(chan error)
	for _, cour := range o.Couriers {
		if len(cour.WorkingHours) == 0 || len(cour.Regions) == 0 {
			return errors.New("error")
		}

		go func() {
			if dublTime(cour.WorkingHours) {
				errChan1 <- errors.New("Duplicate time interval found")
			} else {
				errChan1 <- nil
			}
		}()

		go func() {
			if dublRegion(cour) {
				errChan2 <- errors.New("Duplicate region interval found")
			} else {
				errChan2 <- nil
			}
		}()

		if cour.CourierType != foot && cour.CourierType != bike && cour.CourierType != car {
			return errors.New("error")
		}

		err := <-errChan1
		if err != nil {
			return err
		}
		err = <-errChan2
		if err != nil {
			return err
		}
	}
	return nil
}

func dublRegion(cour dto.CreateCourierDto) bool {
	m := make(map[int32]struct{})
	for _, reg := range cour.Regions {
		if reg < 0 {
			return true
		}
		if _, ok := m[reg]; ok {
			return true
		}
		m[reg] = struct{}{}
	}
	return false
}

func dublTime(cour []string) bool {

	timeRegex := regexp.MustCompile(`^([0-1][0-9]|2[0-3]):[0-5][0-9]-([0-1][0-9]|2[0-3]):[0-5][0-9]$`)
	var intervals []timeInterval

	for _, interval := range cour {
		if err := timeRegex.MatchString(interval); !err {
			return true
		}

		start, end := parseTimeInterval(interval)
		if start > end || start == end {
			return true
		}
		intervals = append(intervals, timeInterval{start, end})
	}

	if len(intervals) > 1 {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].start < intervals[j].start
		})

		for i := 0; i < len(intervals)-1; i++ {
			if intervals[i].end > intervals[i+1].start {
				return true
			}
		}
	}
	return false
}

func parseTimeInterval(interval string) (string, string) {
	parts := regexp.MustCompile(`-`).Split(interval, 2)
	return parts[0], parts[1]
}
