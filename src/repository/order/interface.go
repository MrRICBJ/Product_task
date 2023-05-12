package order

import (
	"context"
	"sss/controllers/dto"
	"sss/entity"
)

type Repo interface {
	GetAll(ctx context.Context, limit, offset int32) ([]entity.Order, error)
	GetById(ctx context.Context, id int64) (*entity.Order, error)
	Create(ctx context.Context, orders *dto.CreateOrderRequest) ([]entity.Order, error)
	Update(ctx context.Context, orders []entity.CompleteOrder) ([]entity.Order, error)
}
