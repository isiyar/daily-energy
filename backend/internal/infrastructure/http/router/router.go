package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/handler"
)

func RegisterRoutes(r gin.IRouter) {
	r.GET("/", handler.HelloWorldHandler)
	r.GET("/ping", handler.PingHandler)
}
