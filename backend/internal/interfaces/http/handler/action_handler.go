package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"github.com/isiyar/daily-energy/backend/pkg/validator"
	"github.com/isiyar/daily-energy/backend/pkg/utils"
	"gorm.io/gorm"
	"log"
	"strconv"
	"net/http"
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
	utgidParam := c.Param("utgid")
	if utgidParam == "" {
		log.Println("Missing utgid in URL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing utgid in URL"})
		return
	}

	utgidInt, err := strconv.ParseInt(utgidParam, 10, 64)
	if err != nil {
		log.Printf("Invalid utgid format in URL: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid in URL"})
		return
	}

	utgidCtx, ok := c.Get("utgid")
	if !ok {
		log.Println("Missing utgid in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing utgid in context"})
		return
	}

	utgidCtxStr, ok := utgidCtx.(string)
	if !ok {
		log.Println("Invalid utgid type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid utgid type in context"})
		return
	}

	utgidCtxInt, err := strconv.ParseInt(utgidCtxStr, 10, 64)
	if err != nil {
		log.Printf("Invalid utgid format in context: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid in context"})
		return
	}

	if utgidInt != utgidCtxInt {
		log.Printf("Utgid mismatch: URL=%d, Context=%d", utgidInt, utgidCtxInt)
		c.JSON(http.StatusForbidden, gin.H{"error": "utgid mismatch"})
		return
	}

	if _, err := h.userUC.Execute(c.Request.Context(), utgidInt); errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
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

	domainAction := req.ToAction(utgidInt)
	if err := h.actionUC.Add(c.Request.Context(), &domainAction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToActionResponse(domainAction))
}

func (h *ActionHandler) GetAction(c *gin.Context) {
	utgidCtx, ok := c.Get("utgid")
	if !ok {
		log.Println("Missing utgid in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing utgid in context"})
		return
	}

	utgidCtxStr, ok := utgidCtx.(string)
	if !ok {
		log.Println("Invalid utgid type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid utgid type in context"})
		return
	}

	utgidCtxInt, err := strconv.ParseInt(utgidCtxStr, 10, 64)
	if err != nil {
		log.Printf("Invalid utgid format in context: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid in context"})
		return
	}

	id := c.Param("id")
	action, err := h.actionUC.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if action.Utgid != utgidCtxInt {
		log.Printf("Utgid mismatch: Action.Utgid=%d, Context=%d", action.Utgid, utgidCtxInt)
		c.JSON(http.StatusForbidden, gin.H{"error": "utgid mismatch"})
		return
	}

	c.JSON(http.StatusOK, dto.ToActionResponse(action))
}

func (h *ActionHandler) GetActions(c *gin.Context) {
	utgidParam := c.Param("utgid")
	if utgidParam == "" {
		log.Println("Missing utgid in URL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing utgid in URL"})
		return
	}

	utgidInt, err := strconv.ParseInt(utgidParam, 10, 64)
	if err != nil {
		log.Printf("Invalid utgid format in URL: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid in URL"})
		return
	}

	utgidCtx, ok := c.Get("utgid")
	if !ok {
		log.Println("Missing utgid in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing utgid in context"})
		return
	}

	utgidCtxStr, ok := utgidCtx.(string)
	if !ok {
		log.Println("Invalid utgid type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid utgid type in context"})
		return
	}

	utgidCtxInt, err := strconv.ParseInt(utgidCtxStr, 10, 64)
	if err != nil {
		log.Printf("Invalid utgid format in context: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid in context"})
		return
	}

	if utgidInt != utgidCtxInt {
		log.Printf("Utgid mismatch: URL=%d, Context=%d", utgidInt, utgidCtxInt)
		c.JSON(http.StatusForbidden, gin.H{"error": "utgid mismatch"})

	startInt, finishInt, err := utils.ParseStartFinish(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if startInt > finishInt {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "start must be less end"})
		return
	}

	if _, err := h.userUC.Execute(c.Request.Context(), utgidInt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	start := c.Query("start_at")
	finish := c.Query("finish_at")

	startInt, err = strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start"})
		return
	}

	finishInt, err = strconv.ParseInt(finish, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end"})
		return
	}

	if startInt > finishInt {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "start must be less than end"})
		return
	}

	actions, err := h.actionUC.GetByStartTimeAndFinishTime(c.Request.Context(), startInt, finishInt, utgidInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToActionsResponse(actions))
}
}