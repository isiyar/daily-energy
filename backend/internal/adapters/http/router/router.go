package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/handler"
)

func RegisterRoutes(r gin.IRouter, h *handler.Handler) {
	r.GET("/ping", handler.PingHandler)

	r.POST("/users", h.User.CreateUser)
	r.GET("/users/:utgid", h.User.GetUser)
	r.PUT("/users/:utgid", h.User.UpdateUser)
	r.DELETE("/users/:utgid", h.User.DeleteUser)

	r.POST("/users/:utgid/actions", h.Action.CreateAction)
	r.GET("/users/:utgid/actions", h.Action.GetActions)
	r.GET("/actions/:id", h.Action.GetAction)
}
