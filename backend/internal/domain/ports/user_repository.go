package ports

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
)

type UserRepository interface {
	GetByUtgid(ctx context.Context, utgid int64) (models.User, error)
	Save(ctx context.Context, user models.User) error
	Delete(ctx context.Context, utgid int64) error
	Update(ctx context.Context, utgid int64, req dto.UserRequest) (models.User, error)
}
