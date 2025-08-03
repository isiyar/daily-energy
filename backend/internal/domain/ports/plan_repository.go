package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type PlanRepository interface {
	GetByStartTimeAndFinishTime(ctx context.Context, startAt, finishAt, utgid int64) ([]models.Plan, error)
	GetByStartTimeAndFinishTimeAndType(ctx context.Context, startAt, finishAt, utgid int64, planType string) ([]models.Plan, error)
	Save(ctx context.Context, plan []models.Plan) error
}
