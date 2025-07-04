package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"github.com/isiyar/daily-energy/backend/pkg/validator"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUC *usecase.UserUseCase
}

func NewUserHandler(userUC *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	utgidStr := c.Param("utgid")

	utgid, err := strconv.ParseInt(utgidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid"})
		return
	}

	user, err := h.userUC.Execute(c.Request.Context(), utgid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(user))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if err := validator.Struct(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	domainUser := req.ToUser()

	if _, err := h.userUC.Execute(c.Request.Context(), domainUser.Utgid); !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	if err := h.userUC.Add(c.Request.Context(), domainUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToUserResponse(domainUser))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	utgidStr := c.Param("utgid")

	utgid, err := strconv.ParseInt(utgidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid utgid"})
		return
	}

	if err := h.userUC.Delete(c.Request.Context(), utgid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
