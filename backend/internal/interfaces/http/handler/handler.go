package handler

type Handler struct {
	Action            *ActionHandler
	User              *UserHandler
	UserWeightHistory *UserWeightHistoryHandler
	Ai                *AiHandler
	Chat              *ChatHandler
}

func NewHandler(actionUC *ActionHandler, userUC *UserHandler, userWeightHistoryUC *UserWeightHistoryHandler, AiUc *AiHandler, ChatUC *ChatHandler) *Handler {
	return &Handler{
		Action:            actionUC,
		User:              userUC,
		UserWeightHistory: userWeightHistoryUC,
		Ai:                AiUc,
		Chat:              ChatUC,
	}
}
