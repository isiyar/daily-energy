package handler

type Handler struct {
	Action *ActionHandler
	User   *UserHandler
}

func NewHandler(actionUC *ActionHandler, userUC *UserHandler) *Handler {
	return &Handler{
		Action: actionUC,
		User:   userUC,
	}
}
