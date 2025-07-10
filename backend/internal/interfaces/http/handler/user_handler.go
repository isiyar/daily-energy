package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/app/usecase"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"github.com/isiyar/daily-energy/backend/pkg/validator"
	"gorm.io/gorm"
	"log"
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

	user, err := h.userUC.Execute(c.Request.Context(), utgidInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(user))
}

func (h *UserHandler) CreateUser(c *gin.Context) {
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
	domainUser.Utgid = utgidInt

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

	if err := h.userUC.Delete(c.Request.Context(), utgidInt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
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

	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if err := validator.Struct(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	newUser, err := h.userUC.Update(c.Request.Context(), utgidInt, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(newUser))
}