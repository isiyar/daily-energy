package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"github.com/isiyar/daily-energy/backend/pkg/utils"
	"github.com/isiyar/daily-energy/backend/pkg/validator"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ActionHandler struct {
	actionUC *usecase.ActionUseCase
	userUC   *usecase.UserUseCase
}

func NewActionHandler(actionUC *usecase.ActionUseCase, userUC *usecase.UserUseCase) *ActionHandler {
	return &ActionHandler{
		actionUC: actionUC,
		userUC:   userUC,
	}
}

func (h *ActionHandler) CreateAction(c *gin.Context) {
	utgid, err := utils.ParseUtgid(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid"})
		return
	}

	if _, err := h.userUC.Execute(c.Request.Context(), utgid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var req dto.ActionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if err := validator.Struct(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	domainAction := req.ToAction(utgid)

	if _, err := h.actionUC.Execute(c.Request.Context(), domainAction.Id); errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	if err := h.actionUC.Add(c.Request.Context(), &domainAction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToActionResponse(domainAction))
}

func (h *ActionHandler) GetAction(c *gin.Context) {
	id := c.Param("id")

	action, err := h.actionUC.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToActionResponse(action))
}

func (h *ActionHandler) GetActions(c *gin.Context) {
	start := c.Query("start_at")
	finish := c.Query("finish_at")

	startInt, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start"})
		return
	}

	finishInt, err := strconv.ParseInt(finish, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end"})
		return
	}

	if startInt > finishInt {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "start must be less end"})
	}

	utgid, err := utils.ParseUtgid(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid"})
		return
	}

	if _, err := h.userUC.Execute(c.Request.Context(), utgid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	actions, err := h.actionUC.GetByStartTimeAndFinishTime(c.Request.Context(), startInt, finishInt, utgid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToActionsResponse(actions))
}
