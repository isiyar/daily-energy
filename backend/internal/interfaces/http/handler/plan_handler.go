package handler

import "github.com/isiyar/daily-energy/backend/internal/app/usecase"

type PlanHandler struct {
	planUC *usecase.PlanUseCase
	userUC *usecase.UserUseCase
}

func NewPlanHandler(planUC *usecase.PlanUseCase, userUC *usecase.UserUseCase) *PlanHandler {
	return &PlanHandler{
		planUC: planUC,
		userUC: userUC,
	}
}
