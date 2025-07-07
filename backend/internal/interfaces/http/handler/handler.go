package handler

type Handler struct {
	Action            *ActionHandler
	User              *UserHandler
  UserWeightHistory *UserWeightHistoryHandler
  Ai                *AiHandler
}

func NewHandler(actionUC *ActionHandler, userUC *UserHandler, userWeightHistoryUC *UserWeightHistoryHandler, AiUc *AiHandler) *Handler {
	return &Handler{
		Action:            actionUC,
		User:              userUC,
    UserWeightHistory: userWeightHistoryUC
    Ai:                AiUc,
  }
}
