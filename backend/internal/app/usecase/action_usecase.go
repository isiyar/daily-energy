package usecase

import "github.com/isiyar/daily-energy/backend/internal/domain/ports"

type ActionUseCase struct {
	repo ports.ActionRepository
}

func NewActionUseCase(repo ports.ActionRepository) *ActionUseCase {
	return &ActionUseCase{repo: repo}
}
