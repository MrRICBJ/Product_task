package service

import (
	"context"
	"sss/internal/controllers/dto"
	"sss/internal/repository/order"
)

type OrderService interface {
	GetOrders(ctx context.Context, limit, offset int32) ([]dto.OrderDto, error)
	GetOrder(ctx context.Context, id int64) (dto.OrderDto, error)
	CreateOrders(ctx context.Context, orders dto.CreateOrderRequest) ([]dto.OrderDto, error)
	CompleteOrders(ctx context.Context, orders dto.CompleteOrderRequestDto) ([]dto.OrderDto, error)
}

type orderService struct {
	repo order.Repo
}

func NewOrderService(repo order.Repo) OrderService {
	return &orderService{repo: repo}
}

func (o *orderService) GetOrders(ctx context.Context, limit, offset int32) ([]dto.OrderDto, error) {
	orders, err := o.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return orders, err
}

func (o *orderService) GetOrder(ctx context.Context, id int64) (dto.OrderDto, error) {
	order, err := o.repo.GetById(ctx, id)
	if err != nil {
		return dto.OrderDto{}, err
	}
	return order, err
}

func (o *orderService) CompleteOrders(ctx context.Context, orders dto.CompleteOrderRequestDto) ([]dto.OrderDto, error) {
	completeOrder, err := o.repo.Update(ctx, orders)
	if err != nil {
		return nil, err
	}
	return completeOrder, err
}

func (o *orderService) CreateOrders(ctx context.Context, orders dto.CreateOrderRequest) ([]dto.OrderDto, error) {
	createOrders, err := o.repo.Create(ctx, orders)
	if err != nil {
		return nil, err
	}
	return createOrders, err
}
