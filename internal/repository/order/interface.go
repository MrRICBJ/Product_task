package order

import (
	"context"
	"sss/internal/controllers/dto"
)

type Repo interface {
	GetAll(ctx context.Context, limit, offset int32) ([]dto.OrderDto, error)
	GetById(ctx context.Context, id int64) (dto.OrderDto, error)
	Create(ctx context.Context, orders dto.CreateOrderRequest) ([]dto.OrderDto, error)
	Update(ctx context.Context, orders dto.CompleteOrderRequestDto) ([]dto.OrderDto, error)
}
