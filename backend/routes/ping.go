package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r gin.IRouter) {
	r.GET("/ping", pingHandler)
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
