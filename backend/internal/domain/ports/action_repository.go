package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type ActionRepository interface {
	GetById(ctx context.Context, id string) (models.Action, error)
	GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishtAt int) (models.Action, error)
	Save(ctx context.Context, user models.Action) error
	Delete(ctx context.Context, id string) error
}
