package usecase

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
)

type UserUseCase struct {
	repo ports.UserRepository
}

func NewUserUseCase(repo ports.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Execute(ctx context.Context, utgid int64) (models.User, error) {
	return uc.repo.GetByUtgid(ctx, utgid)
}

func (uc *UserUseCase) Add(ctx context.Context, user models.User) error {
	return uc.repo.Save(ctx, user)
}

func (uc *UserUseCase) Delete(ctx context.Context, utgid int64) error {
	return uc.repo.Delete(ctx, utgid)
}
