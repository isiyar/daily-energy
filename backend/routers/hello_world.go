package routers

import "github.com/gin-gonic/gin"

func HelloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
