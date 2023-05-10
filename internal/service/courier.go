package service

import (
	"context"
	"sss/internal/controllers/dto"
	"sss/internal/repository/courier"
	"time"
)

type CourService interface {
	GetCouriers(ctx context.Context, limit, offset int32) (*dto.GetCouriersResponse, error)
	CreateCourier(ctx context.Context, cour *dto.CreateCourierRequest) (*dto.CreateCouriersResponse, error)
	GetCourierById(ctx context.Context, id int64) (*dto.CourierDto, error)
	GetCourierMetaInfo(ctx context.Context, id int, startDate, endDate time.Time) (*dto.GetCourierMetaInfoResponse, error)
}

type courService struct {
	repo courier.Repo
}

func NewCourService(repo courier.Repo) CourService {
	return &courService{repo: repo}
}

func (s *courService) GetCouriers(ctx context.Context, limit, offset int32) (*dto.GetCouriersResponse, error) {
	var res dto.GetCouriersResponse
	var err error

	res.Couriers, err = s.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	res.Offset = offset
	res.Limit = limit

	return &res, nil
}

func (s *courService) CreateCourier(ctx context.Context, cour *dto.CreateCourierRequest) (*dto.CreateCouriersResponse, error) {
	res := dto.CreateCouriersResponse{}
	var err error

	res.Couriers, err = s.repo.Create(ctx, cour)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *courService) GetCourierById(ctx context.Context, id int64) (*dto.CourierDto, error) {
	result, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *courService) GetCourierMetaInfo(ctx context.Context, id int, startDate, endDate time.Time) (*dto.GetCourierMetaInfoResponse, error) {
	costs, result, err := s.repo.GetMetaInf(ctx, id, startDate, endDate)
	//if err != nil {
	//	.....
	//}
	result.Rating = calculateRating(startDate, endDate, getCoefficient(result.CourierType), int32(len(costs)))
	result.Earnings = calculateEarnings(costs, getCoefficient(result.CourierType))
	return result, err
}

func calculateRating(startDate, endDate time.Time, c int32, numOrders int32) int32 {
	hours := endDate.Sub(startDate).Hours()
	rating := (numOrders / int32(hours)) * c
	return rating
}

func calculateEarnings(costs []int32, c int32) int32 {
	var earnings int32
	for _, cost := range costs {
		earnings += cost * c
	}
	return earnings
}

func getCoefficient(courierType string) int32 {
	switch courierType {
	case "FOOT":
		return 2
	case "BIKE":
		return 3
	case "CAR":
		return 4
	default:
		return 0
	}
}
