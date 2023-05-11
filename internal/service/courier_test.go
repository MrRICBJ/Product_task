package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"sss/internal/controllers/dto"
	"sss/internal/entity"
	mock_courier "sss/internal/repository/courier/mocks"
	"testing"
	"time"
)

func TestGetCouriers(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_courier.NewMockRepo(ctl)

	ctx := context.Background()

	mockRepo := []entity.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}

	expected := &dto.GetCouriersResponse{
		Couriers: []dto.CourierDto{
			{
				CourierId:    1,
				CourierType:  "FOOT",
				Regions:      []int32{2, 4},
				WorkingHours: []string{"12:34:00", "12:12:00"},
			},
		},
		Limit:  int32(1),
		Offset: int32(0),
	}

	//проверка на валидность
	repo.EXPECT().GetAll(ctx, int32(1), int32(0)).Return(mockRepo, nil).Times(1)

	service := NewCourService(repo)
	orders, err := service.GetCouriers(ctx, int32(1), int32(0))
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().GetAll(ctx, int32(1), int32(0)).Return(nil, errDb).Times(1)

	_, err = service.GetCouriers(ctx, int32(1), int32(0))
	require.Error(t, err)

}

func TestCreateCourier(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_courier.NewMockRepo(ctl)

	ctx := context.Background()

	mockRepo := &dto.CreateCourierRequest{
		Couriers: []dto.CreateCourierDto{
			{
				CourierType:  "FOOT",
				Regions:      []int32{2, 4},
				WorkingHours: []string{"12:34:00", "12:12:00"},
			},
		},
	}

	expected := &dto.CreateCouriersResponse{
		Couriers: []dto.CourierDto{
			{
				CourierId:    1,
				CourierType:  "FOOT",
				Regions:      []int32{2, 4},
				WorkingHours: []string{"12:34:00", "12:12:00"},
			},
		},
	}

	retur := []entity.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}

	//проверка на валидность
	repo.EXPECT().Create(ctx, mockRepo).Return(retur, nil).Times(1)

	service := NewCourService(repo)
	orders, err := service.CreateCourier(ctx, mockRepo)
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().Create(ctx, mockRepo).Return(nil, errDb).Times(1)

	_, err = service.CreateCourier(ctx, mockRepo)
	require.Error(t, err)

}

func TestGetCourierById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_courier.NewMockRepo(ctl)

	ctx := context.Background()

	retur := &entity.Courier{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
	}

	expected := &dto.CourierDto{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
	}

	repo.EXPECT().GetById(ctx, int64(1)).Return(retur, nil).Times(1)

	service := NewCourService(repo)
	orders, err := service.GetCourierById(ctx, int64(1))
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().GetById(ctx, int64(1)).Return(nil, errDb).Times(1)

	_, err = service.GetCourierById(ctx, int64(1))
	require.Error(t, err)
}

func TestGetCourierMetaInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_courier.NewMockRepo(ctl)

	ctx := context.Background()

	retur := []entity.Courier{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
		{
			CourierId:    1,
			CourierType:  "BIKE",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
		{
			CourierId:    1,
			CourierType:  "CAR",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
		{
			CourierId:    1,
			CourierType:  "",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
		},
	}

	expected := []dto.GetCourierMetaInfoResponse{
		{
			CourierId:    1,
			CourierType:  "FOOT",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       0,
			Earnings:     3800,
		},
		{
			CourierId:    1,
			CourierType:  "BIKE",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       0,
			Earnings:     5700,
		},
		{
			CourierId:    1,
			CourierType:  "CAR",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       0,
			Earnings:     7600,
		},
		{
			CourierId:    1,
			CourierType:  "",
			Regions:      []int32{2, 4},
			WorkingHours: []string{"12:34:00", "12:12:00"},
			Rating:       0,
			Earnings:     0,
		},
	}

	cases := []string{
		"test_1", "test_2", "test_3", "test_4",
	}
	t1, _ := time.Parse("2006-01-02", "2023-01-01")
	t2, _ := time.Parse("2006-01-02", "2023-01-02")

	cost := []int32{500, 1000, 400}

	for i, name := range cases {
		t.Run(name, func(t *testing.T) {
			repo.EXPECT().GetMetaInf(ctx, int64(1), t1, t2).Return(cost, &retur[i], nil).Times(1)

			service := NewCourService(repo)
			orders, _ := service.GetCourierMetaInfo(ctx, int64(1), t1, t2)
			require.Equal(t, &expected[i], orders)
		})
	}

	repo.EXPECT().GetMetaInf(ctx, int64(1), time.Time{}, time.Time{}).Return(nil, &retur[0], nil).Times(1)

	service := NewCourService(repo)
	orders, _ := service.GetCourierMetaInfo(ctx, int64(1), time.Time{}, time.Time{})
	newEx := &dto.GetCourierMetaInfoResponse{
		CourierId:    1,
		CourierType:  "FOOT",
		Regions:      []int32{2, 4},
		WorkingHours: []string{"12:34:00", "12:12:00"},
	}
	require.Equal(t, newEx, orders)

}
