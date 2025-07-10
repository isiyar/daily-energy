package usecase

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
)

type PlanUseCase struct {
	repo ports.PlanRepository
}

func NewPlanUseCase(repo ports.PlanRepository) *PlanUseCase {
	return &PlanUseCase{repo: repo}
}

func (uc *PlanUseCase) GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishAt, utgid int64) ([]models.Plan, error) {
	return uc.repo.GetByStartTimeAndFinishTime(ctx, StartAt, FinishAt, utgid)
}

func (uc *PlanUseCase) GetByStartTimeAndFinishTimeAndType(ctx context.Context, StartAt, FinishAt int64, utgid int64, planType string) ([]models.Plan, error) {
	if planType == "" {
		return uc.GetByStartTimeAndFinishTime(ctx, StartAt, FinishAt, utgid)
	}
	return uc.repo.GetByStartTimeAndFinishTimeAndType(ctx, StartAt, FinishAt, utgid, planType)
}

func (uc *PlanUseCase) Add(ctx context.Context, plans []models.Plan) error {
	return uc.repo.Save(ctx, plans)
}
