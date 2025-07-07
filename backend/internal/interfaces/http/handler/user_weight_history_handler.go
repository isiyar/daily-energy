package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"github.com/isiyar/daily-energy/backend/pkg/validator"
)

type UserWeightHistoryHandler struct {
	userWeightHistoryUC *usecase.UserWeightHistoryUseCase
}

func NewUserWeightHistoryHandler(userWeightHistoryUC *usecase.UserWeightHistoryUseCase) *UserWeightHistoryHandler {
	return &UserWeightHistoryHandler{userWeightHistoryUC: userWeightHistoryUC}
}

func (h *UserWeightHistoryHandler) GetUserWeightHistory(c *gin.Context) {
	utgidStr := c.Param("utgid")
	
	utgid, err := strconv.ParseInt(utgidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid"})
		return
	}

	userWeightHistory, err := h.userWeightHistoryUC.GetUserWeightHistory(c.Request.Context(), utgid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(userWeightHistory)

	c.JSON(http.StatusOK, userWeightHistory)

}

func (h *UserWeightHistoryHandler) CreateUserWeightHistory(c *gin.Context) {
	var req dto.UserWeightHistoryCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}
	
	if err := validator.Struct(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}
	
	utgidStr := c.Param("utgid")
	
	utgid, err := strconv.ParseInt(utgidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid"})
		return
	}
	
	domainUserWeightHistory := req.ToUserWeightHistory(utgid)

	if err := h.userWeightHistoryUC.Add(c.Request.Context(), domainUserWeightHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToUserWeightHistoryResponse(domainUserWeightHistory))
}