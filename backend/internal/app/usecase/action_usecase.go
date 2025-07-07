package usecase

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
)

type ActionUseCase struct {
	repo ports.ActionRepository
}

func NewActionUseCase(repo ports.ActionRepository) *ActionUseCase {
	return &ActionUseCase{repo: repo}
}

func (uc *ActionUseCase) Execute(ctx context.Context, id string) (models.Action, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *ActionUseCase) Add(ctx context.Context, action *models.Action) error {
	return uc.repo.Save(ctx, action)
}

func (uc *ActionUseCase) GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishAt, utgid int64) ([]models.Action, error) {
	return uc.repo.GetByStartTimeAndFinishTime(ctx, StartAt, FinishAt, utgid)
}
