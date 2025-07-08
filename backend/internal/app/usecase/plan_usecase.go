package usecase

import (
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
)

type PlanUseCase struct {
	repo ports.PlanRepository
}

func NewPlanUseCase(repo ports.PlanRepository) *PlanUseCase {
	return &PlanUseCase{repo: repo}
}
