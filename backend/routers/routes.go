package routers

import "github.com/gin-gonic/gin"

func RegisterRoutes(r gin.IRouter) {
	r.GET("/", HelloWorldHandler)
	r.GET("/ping", PingHandler)
}
