package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type UserRepository interface {
	GetByUtgid(ctx context.Context, utgid int64) (models.User, error)
	Save(ctx context.Context, user models.User) error
	Delete(ctx context.Context, utgid int64) error
}
