package routers

import "github.com/gin-gonic/gin"

func HelloWorldRouter(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}