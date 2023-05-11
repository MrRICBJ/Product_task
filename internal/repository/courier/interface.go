package courier

import (
	"context"
	"sss/internal/controllers/dto"
	"sss/internal/entity"
	"time"
)

type Repo interface {
	GetAll(ctx context.Context, limit, offset int32) ([]entity.Courier, error)
	GetById(ctx context.Context, id int64) (*entity.Courier, error)
	GetMetaInf(ctx context.Context, id int64, startDate, endDate time.Time) ([]int32, *entity.Courier, error)
	Create(ctx context.Context, cour *dto.CreateCourierRequest) ([]entity.Courier, error)
}
