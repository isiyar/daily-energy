package routers

import (
	"github.com/gin-gonic/gin"
)

func RingRegisterRoutes(r gin.IRouter) {
	r.GET("/ping", PingHandler)
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
