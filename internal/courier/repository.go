package courier

import "context"

type Repository interface {
	GetAll(ctx context.Context, limit, offset int32) (int, interface{})
	GetById(ctx context.Context, id int) (int, interface{})
	Create(ctx context.Context, cour *CreateCourierRequest) (int, interface{})
}
