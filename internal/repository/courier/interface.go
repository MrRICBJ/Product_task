package courier

import (
	"context"
	"sss/internal/controllers/dto"
	"time"
)

type Repo interface {
	GetAll(ctx context.Context, limit, offset int32) ([]dto.CourierDto, error)
	GetById(ctx context.Context, id int64) (*dto.CourierDto, error)
	GetMetaInf(ctx context.Context, id int, startDate, endDate time.Time) ([]int32, *dto.GetCourierMetaInfoResponse, error)
	Create(ctx context.Context, cour *dto.CreateCourierRequest) ([]dto.CourierDto, error)
}
