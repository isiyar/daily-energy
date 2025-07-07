package handler

type Handler struct {
	Action            *ActionHandler
	User              *UserHandler
	UserWeightHistory *UserWeightHistoryHandler
}

func NewHandler(actionUC *ActionHandler, userUC *UserHandler, userWeightHistoryUC *UserWeightHistoryHandler) *Handler {
	return &Handler{
		Action:            actionUC,
		User:              userUC,
		UserWeightHistory: userWeightHistoryUC,
	}
}
