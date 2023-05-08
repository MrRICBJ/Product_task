package courier

import (
	"context"
	"time"
)

type Repository interface {
	GetAll(ctx context.Context, limit, offset int32) (int, interface{})
	GetById(ctx context.Context, id int) (int, interface{})
	GetMetaInf(ctx context.Context, id int, startDate, endDate time.Time) (int, interface{})
	Create(ctx context.Context, cour *CreateCourierRequest) (int, interface{})
}

///couriers/meta-info/{courier_id}
