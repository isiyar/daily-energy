package handler

type Handler struct {
	User *UserHandler
	UserWeightHistory *UserWeightHistoryHandler
}

func NewHandler(userUC *UserHandler, userWeightHistoryUC *UserWeightHistoryHandler) *Handler {
	return &Handler{
		User: userUC,
		UserWeightHistory: userWeightHistoryUC,
	}
}
