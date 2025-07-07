package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type ActionRepository interface {
	GetById(ctx context.Context, id string) (models.Action, error)
	GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishAt, utgid int64) ([]models.Action, error)
	Save(ctx context.Context, action *models.Action) error
}
