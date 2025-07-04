package handler

import "github.com/isiyar/daily-energy/backend/internal/app/usecase"

type ActionHandler struct {
	actionUC *usecase.ActionUseCase
}

func NewActionHandler(actionUC *usecase.ActionUseCase) *ActionHandler {
	return &ActionHandler{actionUC: actionUC}
}
