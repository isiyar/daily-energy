package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/handler"
)

func RegisterRoutes(r gin.IRouter, h *handler.Handler) {
	r.GET("/ping", handler.PingHandler)

	r.POST("/users", h.User.CreateUser)
	r.GET("/users/:utgid", h.User.GetUser)
	r.DELETE("/users/:utgid", h.User.DeleteUser)
}
