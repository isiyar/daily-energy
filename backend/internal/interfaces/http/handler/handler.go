package handler

type Handler struct {
	User *UserHandler
}

func NewHandler(userUC *UserHandler) *Handler {
	return &Handler{User: userUC}
}
