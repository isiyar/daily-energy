package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type UserWeightHistoryRepository interface {
	GetByUtgid(ctx context.Context, utgid int64) (models.UserWeightHistory, error)
	Save(ctx context.Context, user models.UserWeightHistory) error
	Delete(ctx context.Context, utgid int64) error
}
