package courier

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"sss/internal/controllers/dto"
	"sss/internal/entity"
	"time"
)

type repository struct {
	db *sqlx.DB
}

func NewCourRepo(db *sqlx.DB) Repo {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context, limit, offset int32) ([]entity.Courier, error) {
	res := make([]entity.Courier, 0)

	q := `SELECT courier_id, courier_type, regions, working_hours FROM couriers LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := entity.Courier{}
		err = rows.Scan(&tmp.CourierId, &tmp.CourierType, pq.Array(&tmp.Regions), pq.Array(&tmp.WorkingHours))
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (r *repository) GetById(ctx context.Context, id int64) (*entity.Courier, error) {
	cour := entity.Courier{}

	q := `SELECT courier_id, courier_type, regions, working_hours FROM couriers WHERE courier_id = $1`

	err := r.db.QueryRowContext(ctx, q, id).Scan(&cour.CourierId, &cour.CourierType, pq.Array(&cour.Regions), pq.Array(&cour.WorkingHours))
	if err != nil {
		return nil, err
	}

	return &cour, nil
}

func (r *repository) Create(ctx context.Context, cour *dto.CreateCourierRequest) ([]entity.Courier, error) {
	couriers := make([]entity.Courier, 0)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	q := `INSERT INTO couriers (courier_type, regions, working_hours) VALUES ($1, $2, $3) RETURNING courier_id`
	var id int64
	for _, v := range cour.Couriers {
		err := tx.QueryRowContext(ctx, q, v.CourierType, pq.Array(v.Regions), pq.Array(v.WorkingHours)).Scan(&id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		tmp := entity.Courier{}
		tmp.CourierId = id
		tmp.Regions = v.Regions
		tmp.WorkingHours = v.WorkingHours
		tmp.CourierType = v.CourierType
		couriers = append(couriers, tmp)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return couriers, nil
}

func (r *repository) GetMetaInf(ctx context.Context, id int64, startDate, endDate time.Time) ([]int32, *entity.Courier, error) {
	costsList := make([]int32, 0)
	res := entity.Courier{}

	q := `SELECT cost FROM orders WHERE completed_time >= $1 AND completed_time < $2 AND cour_id = $3;`

	_ = r.db.SelectContext(ctx, &costsList, q, startDate, endDate, id)
	//rows, _ := tx.SelectContext(ctx, &costsList, q, startDate, endDate, id)
	//if err != nil {
	//	return nil, nil, err
	//}

	q = `SELECT * FROM couriers WHERE courier_id = $1;`

	_ = r.db.QueryRowContext(ctx, q, id).Scan(&res.CourierId, &res.CourierType, pq.Array(&res.Regions), pq.Array(&res.WorkingHours))

	//if err != nil {
	//	return nil, err
	//}

	return costsList, &res, nil
}
