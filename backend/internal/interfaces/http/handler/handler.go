package handler

type Handler struct {
	Action            *ActionHandler
	User              *UserHandler
	UserWeightHistory *UserWeightHistoryHandler
	Plan              *PlanHandler
	Ai                *AiHandler
}

func NewHandler(actionUC *ActionHandler, userUC *UserHandler, userWeightHistoryUC *UserWeightHistoryHandler, planUC *PlanHandler, AiUc *AiHandler) *Handler {
	return &Handler{
		Action:            actionUC,
		User:              userUC,
		UserWeightHistory: userWeightHistoryUC,
		Plan:              planUC,
		Ai:                AiUc,
	}
}
