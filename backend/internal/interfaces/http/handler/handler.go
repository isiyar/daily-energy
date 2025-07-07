package handler

type Handler struct {
	Action *ActionHandler
	Ai     *AiHandler
	User   *UserHandler
}

func NewHandler(actionUC *ActionHandler, AiUc *AiHandler, userUC *UserHandler) *Handler {
	return &Handler{
		Action: actionUC,
		Ai:     AiUc,
		User:   userUC,
	}
}
