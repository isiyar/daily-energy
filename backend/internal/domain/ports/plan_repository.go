package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type PlanRepository interface {
	GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishtAt int) (models.Plan, error)
	Save(ctx context.Context, user models.Plan) error
	Delete(ctx context.Context, id string) error
}
