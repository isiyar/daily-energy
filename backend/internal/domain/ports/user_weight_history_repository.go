package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type UserWeightHistoryRepository interface {
	GetUserWeightHistory(ctx context.Context, utgid int64) ([]models.UserWeightHistory, error)
	Save(ctx context.Context, userWeightHistory models.UserWeightHistory) error
}
