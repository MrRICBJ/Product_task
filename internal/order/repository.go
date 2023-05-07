package order

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context, limit, offset int32) (int, interface{})
	GetById(ctx context.Context, id int) (int, interface{})
	Create(ctx context.Context, orders *CreateOrderRequest) (int, interface{})
	Update(ctx context.Context, orders *CompleteOrderRequestDto) (int, interface{})
}
