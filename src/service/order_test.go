package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"sss/controllers/dto"
	"sss/entity"
	mock_order "sss/repository/order/mocks"
	"testing"
	time "time"
)

func TestGetOrders(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_order.NewMockRepo(ctl)

	ctx := context.Background()

	expected := []dto.OrderDto{
		{
			OrderId: 1,
			Weight:  10,
			Regions: 5,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
		{
			OrderId: 2,
			Weight:  4,
			Regions: 1,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost: 500,
		},
	}

	mockRepo := []entity.Order{
		{
			OrderId: 1,
			Weight:  10,
			Regions: 5,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
		{
			OrderId: 2,
			Weight:  4,
			Regions: 1,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost: 500,
		},
	}
	//проверка на валидность
	repo.EXPECT().GetAll(ctx, int32(1), int32(0)).Return(mockRepo, nil).Times(1)

	service := NewOrderService(repo)
	orders, err := service.GetOrders(ctx, int32(1), int32(0))
	require.NoError(t, err)
	require.ElementsMatch(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().GetAll(ctx, int32(1), int32(0)).Return(nil, errDb).Times(1)

	service = NewOrderService(repo)
	_, err = service.GetOrders(ctx, int32(1), int32(0))
	require.Error(t, err)
}

func TestGetOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_order.NewMockRepo(ctl)

	ctx := context.Background()

	mockRepo := &entity.Order{
		OrderId: 1,
		Weight:  10,
		Regions: 5,
		DeliveryHours: []string{
			"{10:00-12:00",
			"14:00-16:00}",
		},
		Cost: 500,
	}

	expected := &dto.OrderDto{
		OrderId: 1,
		Weight:  10,
		Regions: 5,
		DeliveryHours: []string{
			"{10:00-12:00",
			"14:00-16:00}",
		},
		Cost: 500,
	}
	//проверка на валидность
	repo.EXPECT().GetById(ctx, int64(1)).Return(mockRepo, nil).Times(1)

	service := NewOrderService(repo)
	orders, err := service.GetOrder(ctx, int64(1))
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().GetById(ctx, int64(1)).Return(nil, errDb).Times(1)

	service = NewOrderService(repo)
	_, err = service.GetOrder(ctx, int64(1))
	require.Error(t, err)
}

func TestCompleteOrders(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_order.NewMockRepo(ctl)

	ctx := context.Background()

	mockRepo := []entity.Order{
		{
			OrderId: 1,
			Weight:  10,
			Regions: 5,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
		{
			OrderId: 2,
			Weight:  4,
			Regions: 1,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
		{
			OrderId: 3,
			Weight:  4,
			Regions: 1,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
	}

	expected := []dto.OrderDto{
		{
			OrderId: 1,
			Weight:  10,
			Regions: 5,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
		{
			OrderId: 2,
			Weight:  4,
			Regions: 1,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
		{
			OrderId: 3,
			Weight:  4,
			Regions: 1,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost:          500,
			CompletedTime: &time.Time{},
		},
	}

	comp := []dto.CompleteOrder{
		{
			OrderId:      1,
			CourierId:    10,
			CompleteTime: time.Time{},
		},
		{
			OrderId:      3,
			CourierId:    10,
			CompleteTime: time.Time{},
		},
		{
			OrderId:      2,
			CourierId:    10,
			CompleteTime: time.Time{},
		},
	}

	var resp = dto.CompleteOrderRequestDto{
		CompleteInfo: comp,
	}
	//проверка на валидность
	repo.EXPECT().Update(ctx, comp).Return(mockRepo, nil).Times(1)

	service := NewOrderService(repo)
	orders, err := service.CompleteOrders(ctx, resp)
	require.NoError(t, err)
	require.ElementsMatch(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().Update(ctx, comp).Return(nil, errDb).Times(1)

	service = NewOrderService(repo)
	_, err = service.CompleteOrders(ctx, resp)
	require.Error(t, err)
}

func TestCreateOrders(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_order.NewMockRepo(ctl)

	ctx := context.Background()

	createOrder := []dto.CreateOrderDto{
		{
			Weight:  34,
			Regions: 5,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost: 500,
		},
	}
	comp := &dto.CreateOrderRequest{
		Orders: createOrder,
	}

	mockRepo := []entity.Order{
		{
			OrderId: 1,
			Weight:  34,
			Regions: 5,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost: 500,
		},
	}

	expected := []dto.OrderDto{
		{
			OrderId: 1,
			Weight:  34,
			Regions: 5,
			DeliveryHours: []string{
				"{10:00-12:00",
				"14:00-16:00}",
			},
			Cost: 500,
		},
	}

	repo.EXPECT().Create(ctx, comp).Return(mockRepo, nil).Times(1)

	service := NewOrderService(repo)
	orders, err := service.CreateOrders(ctx, comp)
	require.NoError(t, err)
	require.ElementsMatch(t, expected, orders)

	//проверка на невалидность
	errDb := errors.New("db is down")
	repo.EXPECT().Create(ctx, comp).Return(nil, errDb).Times(1)

	service = NewOrderService(repo)
	_, err = service.CreateOrders(ctx, comp)
	require.Error(t, err)
}
